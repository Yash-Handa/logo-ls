// this file contain dir type definition
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"text/tabwriter"
	"time"
)

type file struct {
	name, ext, indicator string
	modTime              time.Time
	size                 int64 // in bytes
	mode                 string
	modeBits             uint32
	owner, group         string // use syscall package
	blocks               int64  // blocks required by the file multiply buy 512 to get block size
}

type dir struct {
	info  *file
	files []*file // all child files and dirs
	dirs  []*file // for recursion contain only child dirs
	less  func(int, int) bool
}

// define methods on *dir type only not on file type

func newDir(d *os.File) (*dir, error) {
	// some flag variable combinations
	var long, curDir, showHidden bool = flagVector&(flag_l|flag_o|flag_g) > 0, flagVector&(flag_a) > 0, flagVector&(flag_a|flag_A) > 0

	t := new(dir)

	// filing current dir info
	t.info = new(file)
	t.info.name = d.Name()
	t.info.ext = ""
	t.info.indicator = "/"
	if curDir {
		ds, err := d.Stat()
		if err != nil {
			return nil, err
		}
		t.info.size = ds.Size()
		t.info.modTime = ds.ModTime()
		if long {
			t.info.mode = ds.Mode().String()
			t.info.modeBits = uint32(ds.Mode())
			t.info.owner, t.info.group = getOwnerGroupInfo(ds)
		}
		if flagVector&flag_s > 0 {
			if s, ok := ds.Sys().(*syscall.Stat_t); ok {
				t.info.blocks = s.Blocks
			}
		}
	}

	files, err := d.Readdir(0)
	for _, v := range files {
		name := v.Name()
		if !showHidden && strings.HasPrefix(name, ".") {
			continue
		}

		// don't fill files info if the -d flag is passed
		if flagVector&flag_d > 0 && v.IsDir() == false {
			continue
		}

		f := new(file)
		f.ext = filepath.Ext(name)
		f.name = name[0 : len(name)-len(f.ext)]
		f.indicator = getIndicator(v.Mode())
		f.size = v.Size()
		f.modTime = v.ModTime()
		if long {
			f.mode = v.Mode().String()
			f.modeBits = uint32(v.Mode())
			f.owner, f.group = getOwnerGroupInfo(v)
		}
		if flagVector&flag_s > 0 {
			if s, ok := v.Sys().(*syscall.Stat_t); ok {
				f.blocks = s.Blocks
			}
		}
		if flagVector&flag_d == 0 {
			t.files = append(t.files, f)
		}
		if v.IsDir() {
			t.dirs = append(t.dirs, f)
		}
	}

	// return *dir with no error
	// or partial *dir with error (produced by Readdir)
	return t, err
}

func (d *dir) print() *bytes.Buffer {
	// take care of printing, extending symbolic links in long forms

	buf := bytes.NewBuffer([]byte(""))
	var w *tabwriter.Writer
	switch {
	case flagVector&(flag_l|flag_o|flag_g) > 0:
		w = tabwriter.NewWriter(buf, 0, 0, 1, ' ', tabwriter.DiscardEmptyColumns)
		fmtStr := "%s\t%s\t%s\t%s\t%s\t%s\t\n"
		for _, v := range d.files {
			if flagVector&flag_s > 0 {
				fmt.Fprintf(w, "%s\t", getSizeInFormate(v.blocks*512))
			}
			fmt.Fprintf(w, fmtStr, v.mode, v.owner, v.group, getSizeInFormate(v.size), v.modTime.Format(time.Stamp), v.name+v.ext+v.indicator)
		}
	case flagVector&flag_1 > 0:
		w = tabwriter.NewWriter(buf, 0, 0, 1, ' ', tabwriter.DiscardEmptyColumns)
		for _, v := range d.files {
			if flagVector&flag_s > 0 {
				fmt.Fprintf(w, "%s\t", getSizeInFormate(v.blocks*512))
			}
			fmt.Fprintf(w, "%s\t\n", v.name+v.ext+v.indicator)
		}
	default:
		w = tabwriter.NewWriter(buf, 0, 0, 2, ' ', tabwriter.DiscardEmptyColumns)
		for _, v := range d.files {
			s := ""
			if flagVector&flag_s > 0 {
				s = getSizeInFormate(v.blocks*512) + " "
			}
			fmt.Fprintf(w, "%s\t", s+v.name+v.ext+v.indicator)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
	return buf
}

// sorting functions
func (d *dir) Len() int {
	return len(d.files)
}

func (d *dir) Swap(i, j int) {
	d.files[i], d.files[j] = d.files[j], d.files[i]
}

func (d *dir) Less(i, j int) bool {
	return d.less(i, j)
}

// Custom less functions
func lessFuncGenerator(d *dir) func(int, int) bool {
	switch {
	case (flagVector & flag_alpha) > 0:
		// sort by alphabetical order of name.ext
		return func(i, j int) bool {
			return d.files[i].name+d.files[i].ext < d.files[j].name+d.files[j].ext
		}
	case (flagVector & flag_S) > 0:
		// sort by file size, largest first
		return func(i, j int) bool {
			return d.files[i].size > d.files[j].size
		}
	case (flagVector & flag_t) > 0:
		// sort by modification time, newest first
		return func(i, j int) bool {
			return d.files[i].modTime.Before(d.files[j].modTime)
		}
	case (flagVector & flag_X) > 0:
		// sort alphabetically by entry extension
		return func(i, j int) bool {
			return d.files[i].ext < d.files[j].ext
		}
	case (flagVector & flag_v) > 0:
		// natural sort of (version) numbers within text
		return func(i, j int) bool {
			return d.files[i].name+d.files[i].ext < d.files[j].name+d.files[j].ext
		}
	default:
		return nil
	}
}

// get Owner and Group info
var grpMap = make(map[string]string)
var userMap = make(map[string]string)

func getOwnerGroupInfo(fi os.FileInfo) (o string, g string) {
	if stat, ok := fi.Sys().(*syscall.Stat_t); ok {
		if flagVector&(flag_l|flag_o) > 0 {
			UID := strconv.Itoa(int(stat.Uid))
			if n, ok := userMap[UID]; ok {
				o = n
			} else {
				u, err := user.LookupId(UID)
				if err != nil {
					o = ""
				} else {
					o = u.Name
					userMap[UID] = u.Name
				}
			}
		}

		if flagVector&flag_G == 0 && flagVector&(flag_l|flag_g) > 0 {
			GID := strconv.Itoa(int(stat.Gid))
			if n, ok := grpMap[GID]; ok {
				g = n
			} else {
				grp, err := user.LookupGroupId(GID)
				if err != nil {
					g = ""
				} else {
					g = grp.Name
					grpMap[GID] = grp.Name
				}
			}
		}
	}

	return
}

// get indicator of the file
func getIndicator(modebit os.FileMode) (i string) {
	switch {
	case modebit&os.ModeDir > 0:
		i = "/"
	case modebit&1000000 > 0:
		i = "*"
	case modebit&os.ModeNamedPipe > 0:
		i = "|"
	case modebit&os.ModeSymlink > 0:
		i = "@"
	case modebit&os.ModeSocket > 0:
		i = "="
	}
	return i
}

func getSizeInFormate(b int64) string {
	if flagVector&flag_h == 0 {
		return fmt.Sprintf("%d", b)
	}

	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c",
		float64(b)/float64(div), "KMGTPE"[exp])
}
