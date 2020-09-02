// this file contain dir type definition
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"text/tabwriter"
	"time"

	"github.com/Yash-Handa/logo-ls/ctw"
)

type file struct {
	name, ext, indicator string
	modTime              time.Time
	size                 int64 // in bytes
	mode                 string
	modeBits             uint32
	owner, group         string // use syscall package
	blocks               int64  // blocks required by the file multiply buy 512 to get block size
	// 'U'-> untracked file 'M'-> Modified file '●'-> modified dir ' '-> Not Updated/ not in git repo
	gitStatus rune
}

type dir struct {
	info   *file
	parent *file
	files  []*file  // all child files and dirs
	dirs   []string // for recursion contain only child dirs
	less   func(int, int) bool
}

// define methods on *dir type only not on file type

func newDir(d *os.File) (*dir, error) {
	// some flag variable combinations
	var long, curDir, showHidden bool = flagVector&(flag_l|flag_o|flag_g) > 0, flagVector&(flag_a|flag_d) > 0, flagVector&(flag_a|flag_A) > 0

	t := new(dir)

	// filing current dir info
	t.info = new(file)
	t.info.name = "."
	ds, err := d.Stat()
	if err != nil {
		return nil, err
	}

	// getting Git Status of the entire repository
	var gitRepoStatus map[string]rune // could be nil
	if flagVector&flag_D == 0 {
		gitRepoStatus = getFilesGitStatus(d.Name()) // returns map or nil
	}

	if curDir {
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

		if gitRepoStatus != nil {
			if v.IsDir() {
				f.gitStatus = gitRepoStatus[v.Name()+"/"]
			} else {
				f.gitStatus = gitRepoStatus[v.Name()]
			}
		}

		t.files = append(t.files, f)
		if v.IsDir() {
			t.dirs = append(t.dirs, name+"/")
		}
	}

	// if -a flag is passed then only eval parent dir and append to files
	if flagVector&flag_a > 0 {
		t.files = append(t.files, t.info)
		p, err := filepath.Abs(d.Name())
		if err != nil {
			// partial *dir (without parent dir) and error
			return t, err
		}
		pp := filepath.Dir(p)
		pds, err := os.Lstat(pp)
		if err != nil {
			// partial *dir (without parent dir) and error
			return t, err
		}
		t.parent = new(file)
		t.parent.name = ".."
		t.parent.size = pds.Size()
		t.parent.modTime = pds.ModTime()
		if long {
			t.parent.mode = pds.Mode().String()
			t.parent.modeBits = uint32(pds.Mode())
			t.parent.owner, t.parent.group = getOwnerGroupInfo(pds)
		}
		if flagVector&flag_s > 0 {
			if s, ok := pds.Sys().(*syscall.Stat_t); ok {
				t.parent.blocks = s.Blocks
			}
		}
		t.files = append(t.files, t.parent)
	}

	// return *dir with no error
	// or partial *dir with error (produced by Readdir)
	return t, err
}

func newDir_ArgFiles(files []os.FileInfo) *dir {
	var long bool = flagVector&(flag_l|flag_o|flag_g) > 0

	t := new(dir)

	for _, v := range files {
		name := v.Name()
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
	}
	return t
}

func newDirs_Recussion(d *os.File) {
	dd, err := newDir(d)
	d.Close()
	if err != nil {
		log.Printf("partial access to %q: %v\n", d.Name(), err)
		_ = set_osExitCode(code_Minor)
	}
	// print the info of the files of the directory
	io.Copy(os.Stdout, dd.print())
	if len(dd.dirs) == 0 {
		return
	}
	// at this point dd.print has sorted the children files
	// but not using it instead printing children in directory order
	temp := make([]string, len(dd.dirs))
	for i, v := range dd.dirs {
		temp[i] = filepath.Join(d.Name(), v)
	}
	for _, v := range temp {
		fmt.Printf("\n%s:\n", v)
		f, err := os.Open(v)
		if err != nil {
			log.Printf("cannot access %q: %v\n", v, err)
			f.Close()
			_ = set_osExitCode(code_Minor)
			continue
		}
		newDirs_Recussion(f)
	}
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
			fmt.Fprintf(w, fmtStr, v.mode, v.owner, v.group, getSizeInFormate(v.size), v.modTime.Format(timeFormate), v.name+v.ext+v.indicator)
		}
		w.Flush()
	case flagVector&flag_1 > 0:
		w = tabwriter.NewWriter(buf, 0, 0, 1, ' ', tabwriter.DiscardEmptyColumns)
		for _, v := range d.files {
			if flagVector&flag_s > 0 {
				fmt.Fprintf(w, "%s\t", getSizeInFormate(v.blocks*512))
			}
			fmt.Fprintf(w, "%s\t\n", v.name+v.ext+v.indicator)
		}
		w.Flush()
	default:
		// w = tabwriter.NewWriter(buf, 0, 0, 2, ' ', tabwriter.DiscardEmptyColumns)
		temp := make([]string, len(d.files))
		tempG := make([]string, len(d.files))
		noGit := true
		for i, v := range d.files {
			s := ""
			if flagVector&flag_s > 0 {
				s = getSizeInFormate(v.blocks*512) + " "
			}
			temp[i] = s + v.name + v.ext + v.indicator

			if v.gitStatus != '\x00' && v.gitStatus != ' ' {
				tempG[i] = " " + string(v.gitStatus)
				noGit = false
			}
		}
		if noGit {
			ctw.Ctw(buf, temp, terminalWidth)
		} else {
			ctw.CtwGit(buf, temp, tempG, terminalWidth)
		}
	}
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
