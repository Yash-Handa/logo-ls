// this file contain dir type definition
package dir

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/Yash-Handa/logo-ls/assets"
	"github.com/Yash-Handa/logo-ls/internal/api"
	"github.com/Yash-Handa/logo-ls/internal/ctw"
	"github.com/Yash-Handa/logo-ls/internal/sysState"
	"github.com/mattn/go-colorable"
)

// create the open dir icon
var OpenDirIcon = assets.Icon_Def["diropen"].GetColor(1) + assets.Icon_Def["diropen"].GetGlyph() + "\033[0m" + " "

type file struct {
	name, ext, indicator string
	modTime              time.Time
	size                 int64 // in bytes
	mode                 string
	modeBits             uint32
	owner, group         string // use syscall package
	blocks               int64  // blocks required by the file multiply buy 512 to get block size
	// 'U'-> untracked file 'M'-> Modified file '●'-> modified dir ' '-> Not Updated/ not in git repo
	gitStatus string
	icon      string
	iconColor string
}

type dir struct {
	info   *file
	parent *file
	files  []*file  // all child files and dirs
	dirs   []string // for recursion contain only child dirs
	less   func(int, int) bool
}

// define methods on *dir type only not on file type

func New(d *os.File) (*dir, error) {
	// some flag variable combinations
	var long, curDir, showHidden bool = api.FlagVector&(api.Flag_l|api.Flag_o|api.Flag_g) > 0, api.FlagVector&(api.Flag_a|api.Flag_d) > 0, api.FlagVector&(api.Flag_a|api.Flag_A) > 0

	t := new(dir)

	// filing current dir info
	t.info = new(file)
	t.info.name = "."
	ds, err := d.Stat()
	if err != nil {
		return nil, err
	}

	// getting Git Status of the entire repository
	var gitRepoStatus map[string]string // could be nil
	if api.FlagVector&api.Flag_D > 0 {
		gitRepoStatus = getFilesGitStatus(d.Name()) // returns map or nil
		if len(gitRepoStatus) == 0 {
			gitRepoStatus = nil
		}
	}

	if curDir {
		t.info.size = ds.Size()
		t.info.modTime = ds.ModTime()
		if long {
			t.info.mode = ds.Mode().String()
			t.info.modeBits = uint32(ds.Mode())
			t.info.owner, t.info.group = getOwnerGroupInfo(ds)
		}
		if api.FlagVector&api.Flag_s > 0 {
			dirBlocks(t.info, ds)
		}
		if api.FlagVector&api.Flag_i == 0 {
			t.info.icon = assets.Icon_Def["diropen"].GetGlyph()
			if api.FlagVector&api.Flag_c == 0 {
				t.info.iconColor = assets.Icon_Def["diropen"].GetColor(1)
			}
		}
	}

	// don't fill files info if the -d flag is passed
	if api.FlagVector&api.Flag_d > 0 {
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
		if api.FlagVector&api.Flag_s > 0 {
			dirBlocks(f, v)
		}

		if api.FlagVector&api.Flag_i == 0 {
			f.icon, f.iconColor = getIcon(f.name, f.ext, f.indicator)
			if api.FlagVector&api.Flag_c != 0 {
				f.iconColor = ""
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
	if api.FlagVector&api.Flag_a > 0 {
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
		if api.FlagVector&api.Flag_s > 0 {
			dirBlocks(t.parent, pds)
		}
		if api.FlagVector&api.Flag_i == 0 {
			t.parent.icon = assets.Icon_Def["diropen"].GetGlyph()
			if api.FlagVector&api.Flag_c == 0 {
				t.parent.iconColor = assets.Icon_Def["diropen"].GetColor(1)
			}
		}
		t.files = append(t.files, t.parent)
	}

	// return *dir with no error
	// or partial *dir with error (produced by Readdir)
	return t, err
}

func New_ArgFiles(files []os.FileInfo) *dir {
	var long bool = api.FlagVector&(api.Flag_l|api.Flag_o|api.Flag_g) > 0

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
		if api.FlagVector&api.Flag_s > 0 {
			dirBlocks(f, v)
		}
		if api.FlagVector&api.Flag_i == 0 {
			f.icon, f.iconColor = getIcon(f.name, f.ext, f.indicator)
			if api.FlagVector&api.Flag_c != 0 {
				f.iconColor = ""
			}
		}
		t.files = append(t.files, f)
	}
	return t
}

func New_Recussion(d *os.File) {
	dd, err := New(d)
	d.Close()
	if err != nil {
		log.Printf("partial access to %q: %v\n", d.Name(), err)
		sysState.ExitCode(sysState.Code_Minor)
	}
	// print the info of the files of the directory
	var out io.Writer = os.Stdout
	if runtime.GOOS == "windows" {
		out = colorable.NewColorableStdout()
	}
	io.Copy(out, dd.Print())
	if len(dd.dirs) == 0 {
		return
	}
	// at this point dd.print has sorted the children files
	// but not using it instead printing children in directory order
	temp := make([]string, len(dd.dirs))
	sort.Strings(dd.dirs)
	for i, v := range dd.dirs {
		temp[i] = filepath.Join(d.Name(), v)
	}
	for _, v := range temp {
		fmt.Printf("\n%s:\n", OpenDirIcon+v)
		f, err := os.Open(v)
		if err != nil {
			log.Printf("cannot access %q: %v\n", v, err)
			f.Close()
			sysState.ExitCode(sysState.Code_Minor)
			continue
		}
		New_Recussion(f)
	}
}

func (d *dir) Print() *bytes.Buffer {
	// take care of printing, extending symbolic links in long forms

	//sorting
	lessFuncGenerator(d)
	if api.FlagVector&api.Flag_U == 0 && api.FlagVector&api.Flag_r > 0 {
		sort.Sort(sort.Reverse(d))
	} else {
		sort.Sort(d)
	}

	buf := bytes.NewBuffer([]byte(""))
	switch {
	case api.FlagVector&(api.Flag_l|api.Flag_o|api.Flag_g) > 0:
		w := ctw.NewLong(9)
		for _, v := range d.files {
			if api.FlagVector&api.Flag_s > 0 {
				w.AddRow(getSizeInFormate(v.blocks*512), v.mode, v.owner, v.group, getSizeInFormate(v.size), v.modTime.Format(api.GetTimeFormate()), v.icon, v.name+v.ext+v.indicator, v.gitStatus)
			} else {
				w.AddRow("", v.mode, v.owner, v.group, getSizeInFormate(v.size), v.modTime.Format(api.GetTimeFormate()), v.icon, v.name+v.ext+v.indicator, v.gitStatus)
			}
			w.IconColor(v.iconColor)
		}
		w.Flush(buf)
	case api.FlagVector&api.Flag_1 > 0:
		w := ctw.NewLong(4)
		for _, v := range d.files {
			if api.FlagVector&api.Flag_s > 0 {
				w.AddRow(getSizeInFormate(v.blocks*512), v.icon, v.name+v.ext+v.indicator, v.gitStatus)
			} else {
				w.AddRow("", v.icon, v.name+v.ext+v.indicator, v.gitStatus)
			}
			w.IconColor(v.iconColor)
		}
		w.Flush(buf)
	default:
		w := ctw.New(sysState.GetTerminalWidth())
		for _, v := range d.files {
			if api.FlagVector&api.Flag_s > 0 {
				w.AddRow(getSizeInFormate(v.blocks*512), v.icon, v.name+v.ext+v.indicator, v.gitStatus)
			} else {
				w.AddRow("", v.icon, v.name+v.ext+v.indicator, v.gitStatus)
			}
			w.IconColor(v.iconColor)
		}
		w.Flush(buf)
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
