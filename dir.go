// this file contain dir type definition
package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
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
	var long, curDir, showHidden bool = flagVector&(flag_l|flag_o|flag_g) > 0, flagVector&(flag_a|flag_d) > 0, flagVector&(flag_a|flag_A) > 0

	t := new(dir)

	// filing current dir info
	t.info = new(file)
	t.info.name = d.Name()
	t.info.ext = ""
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

	// don't fill files info if the -d flag is passed
	if flagVector&flag_d > 0 {
		t.files = append(t.files, t.info)
		return t, nil
	}

	files, err := d.Readdir(0)
	for _, v := range files {
		name := v.Name()
		if !showHidden && strings.HasPrefix(name, ".") {
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
		t.files = append(t.files, f)
		if v.IsDir() {
			t.dirs = append(t.dirs, f)
		}
	}

	if flagVector&flag_a > 0 {
		t.files = append(t.files, t.info)
	}

	// return *dir with no error
	// or partial *dir with error (produced by Readdir)
	return t, err
}

func (d *dir) print() *bytes.Buffer {
	// take care of printing, extending symbolic links in long forms

	//sorting
	lessFuncGenerator(d)
	if flagVector&flag_U == 0 && flagVector&flag_r > 0 {
		sort.Sort(sort.Reverse(d))
	} else {
		sort.Sort(d)
	}

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
