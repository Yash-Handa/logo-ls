package assets_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"testing"

	"github.com/Yash-Handa/logo-ls/assets"
	"github.com/Yash-Handa/logo-ls/internal/ctw"
	"golang.org/x/crypto/ssh/terminal"
)

func TestFileIcons(t *testing.T) {
	log.Println("Printing each supported file name and ext by the icon pack")

	// get terminal width
	terminalWidth, _, e := terminal.GetSize(int(os.Stdout.Fd()))
	if e != nil {
		terminalWidth = 80
	}

	ks := make([]string, 0)
	for k := range assets.Icon_Set {
		ks = append(ks, k)
	}
	sort.Strings(ks)

	for _, v := range ks {
		t.Run("Testing icon: "+v, func(st *testing.T) {
			i := assets.Icon_Set[v]
			fmt.Fprintln(os.Stderr)
			buf := bytes.NewBuffer([]byte(""))
			log.Println("Printing files of type", i.GetColor(1)+v+"\033[0m")
			w := ctw.New(terminalWidth)
			for f, d := range assets.Icon_FileName {
				if d == i {
					w.AddRow("    ", d.GetGlyph(), f, "")
					w.IconColor(d.GetColor(1))
				}
			}
			w.Flush(buf)
			io.Copy(os.Stderr, buf)

			buf = bytes.NewBuffer([]byte(""))
			log.Println("Printing extentions of type", i.GetColor(1)+v+"\033[0m")
			w = ctw.New(terminalWidth)
			for e, d := range assets.Icon_Ext {
				if d == i {
					w.AddRow("    ", d.GetGlyph(), e, "")
					w.IconColor(d.GetColor(1))
				}
			}
			w.Flush(buf)
			io.Copy(os.Stderr, buf)
		})
	}
}

func TestIconDisplay(t *testing.T) {
	// get terminal width
	terminalWidth, _, e := terminal.GetSize(int(os.Stdout.Fd()))
	if e != nil {
		terminalWidth = 80
	}

	temp := [2]map[string]*assets.Icon_Info{assets.Icon_Set, assets.Icon_Def}

	for i, set := range temp {
		t.Run(fmt.Sprintf("Icon Set %d", i+1), func(st *testing.T) {
			//sorting alphabetically
			ks := make([]string, 0)
			for k := range set {
				ks = append(ks, k)
			}
			sort.Strings(ks)

			// display icons
			buf := bytes.NewBuffer([]byte("\n"))
			w := ctw.New(terminalWidth)
			for _, v := range ks {
				w.AddRow("    ", set[v].GetGlyph(), v, "")
				w.IconColor(set[v].GetColor(1))
			}
			w.Flush(buf)
			io.Copy(os.Stdout, buf)
			fmt.Fprintln(os.Stdout)
		})
	}
}
