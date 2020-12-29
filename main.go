package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sort"

	"github.com/Yash-Handa/logo-ls/assets"
	"github.com/Yash-Handa/logo-ls/internal/api"
	"github.com/Yash-Handa/logo-ls/internal/ctw"
	"github.com/Yash-Handa/logo-ls/internal/dir"
	"github.com/Yash-Handa/logo-ls/internal/sysState"
	"github.com/mattn/go-colorable"
)

func main() {
	// load flags into api.FlagVector and other sysState update stuff
	api.Bootstrap()

	// config ctw and OpenDirIcon
	if api.FlagVector&api.Flag_c > 0 {
		ctw.DisplayColor(false)
		dir.OpenDirIcon = assets.Icon_Def["diropen"].GetGlyph() + " "
	}
	if api.FlagVector&api.Flag_i > 0 {
		dir.OpenDirIcon = ""
		ctw.DisplayBrailEmpty(false)
	}

	// extract files/dir from arguments
	dirs := api.FileList
	if len(dirs) == 0 {
		// use pwd
		dirs = append(dirs, ".")
	}

	sort.Strings(dirs)

	args := struct {
		files []os.FileInfo
		dirs  []*os.File
	}{}

	// segregate args in files and dirs, and print error for those which cannot be opened
	for _, v := range dirs {
		d, err := os.Open(v)
		if err != nil {
			log.Printf("cannot access %q: %v\n", v, err)
			d.Close()
			sysState.ExitCode(sysState.Code_Serious)
			continue
		}
		ds, err := d.Stat()
		if err != nil {
			log.Printf("cannot access %q: %v\n", v, err)
			d.Close()
			sysState.ExitCode(sysState.Code_Serious)
			continue
		}
		if ds.IsDir() {
			args.dirs = append(args.dirs, d)
		} else {
			args.files = append(args.files, ds)
		}
	}

	var out io.Writer = os.Stdout
	if runtime.GOOS == "windows" {
		out = colorable.NewColorableStdout()
	}

	// process and display all files
	if len(args.files) > 0 {
		io.Copy(out, dir.New_ArgFiles(args.files).Print())
		if len(args.dirs) > 0 {
			fmt.Println()
		}
	}

	// process and display all the dirs in arg
	if api.FlagVector&api.Flag_R > 0 {
		// use recursive func
		for i, v := range args.dirs {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("%s:\n", dir.OpenDirIcon+v.Name())
			if api.FlagVector&api.Flag_D > 0 {
				dir.GitRepoCompute()
			}
			dir.New_Recussion(v)
		}
	} else {
		pName := len(dirs) > 1
		for i, v := range args.dirs {
			if pName {
				fmt.Printf("%s:\n", dir.OpenDirIcon+v.Name())
			}
			if api.FlagVector&api.Flag_D > 0 {
				dir.GitRepoCompute()
			}
			d, err := dir.New(v)
			v.Close()
			if err != nil {
				log.Printf("partial access to %q: %v\n", v.Name(), err)
				sysState.ExitCode(sysState.Code_Serious)
			}
			// print the info of the files of the directory
			io.Copy(out, d.Print())
			if i < len(args.dirs)-1 {
				fmt.Println()
			}
		}
	}
	os.Exit(sysState.GetExitCode())
}

func init() {
	log.SetPrefix("logo-ls: ")
	log.SetFlags(0)
}
