// this file contain dir type definition
package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

type file struct {
	name, ext, indicator string
	modTime              time.Time
	size                 int64 // in bytes
	mode                 string
	modeBits             uint32
	owner, group         string // use syscall package
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
		t.files = append(t.files, f)
		if v.IsDir() {
			t.dirs = append(t.dirs, f)
		}
	}

	// return *dir with no error
	// or partial *dir with error (produced by Readdir)
	return t, err
}

func (d *dir) print() []byte {
	// take care of printing, extending symbolic links in long forms

	// a dummy print:
	var t []byte
	t = append(t, "Name of the dir: "+d.info.name+d.info.ext+d.info.indicator+"\n"...)
	for _, v := range d.files {
		t = append(t, v.name+v.ext+v.indicator+"\t"...)
	}
	return t
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
func getOwnerGroupInfo(fi os.FileInfo) (o string, g string) {
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
