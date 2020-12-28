package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

type nodes []struct {
	name        string
	textContent []byte
	dirContent  *nodes
}

func (n *nodes) create(t *testing.T, dir string) {
	for _, f := range *n {
		fn := filepath.Join(dir, f.name)

		if f.textContent != nil {
			if err := ioutil.WriteFile(fn, f.textContent, 0666); err != nil {
				t.Fatal(err)
			}
		} else if f.dirContent != nil {
			if err := os.Mkdir(fn, 0777); err != nil {
				t.Log(err)
				t.Fail()
				continue
			}
			f.dirContent.create(t, fn)
		}
	}
}

func TestE2E(t *testing.T) {
	goExec, err := exec.LookPath("go")
	if err != nil {
		t.Fatal(err)
	}

	dirEnv := filepath.Join("./testdata", "dirEnv")
	os.RemoveAll(dirEnv)

	// create a temp directory environment
	if err := os.Mkdir(dirEnv, 0777); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dirEnv)

	// if you update the fileStruct or create any changes that may break the snapshorts present in ./testdata/logo-ls*.txt
	// then you do have to update the snapshorts like: logo-ls -V > ./testdata/logo-ls-V.txt
	fileStruct := nodes{
		{name: "1test.go", textContent: []byte("package main\nimport \"fmt\"\nfunc main() {\n  fmt.Println(\"Hello\")\n}")},
		{name: "2test.js", textContent: []byte("consol.log('Hello')")},
		{name: "test.routes.tsx", textContent: []byte("consol.log('Hello')")},
		{name: "Dockerfile", textContent: []byte("FROM golang:1.15.2-buster AS base")},
		{name: "abc.css", textContent: []byte("h1 {\n  color : red;\n}")},
		{name: "abc.sass", textContent: []byte("h1 {\n  color : red;\n}")},
		{name: ".abc.txt", textContent: []byte("Hello")},
		{name: "3test.py", textContent: []byte("print('Hello')")},
		{name: "testDir", dirContent: &nodes{
			{name: "abc.txt", textContent: []byte("Hello")},
			{name: "1test.go", textContent: []byte("package main\nimport \"fmt\"\nfunc main() {\n  fmt.Println(\"Hello\")\n}")},
			{name: "2test.js", textContent: []byte("consol.log('Hello')")},
		}},
		{name: ".privateDir", dirContent: &nodes{
			{name: "abc.txt", textContent: []byte("Hello")},
			{name: "1test.go", textContent: []byte("package main\nimport \"fmt\"\nfunc main() {\n  fmt.Println(\"Hello\")\n}")},
			{name: "2test.js", textContent: []byte("consol.log('Hello')")},
		}},
		{name: "Downloads", dirContent: &nodes{}},
	}
	fileStruct.create(t, dirEnv)

	tt := []struct {
		args     []string
		testFile string
		td       string
	}{
		{args: []string{"-1"}, testFile: "logo-ls.snap", td: "Testing normal execution"},
		{args: []string{"-1a"}, testFile: "logo-ls-a.snap", td: "Testing -a (all) execution"},
		{args: []string{"-1A"}, testFile: "logo-ls-A.snap", td: "Testing -A (almost all) execution"},
		{args: []string{"-1i"}, testFile: "logo-ls-i.snap", td: "Testing -i (no icon) execution"},
		{args: []string{"-1r"}, testFile: "logo-ls-r.snap", td: "Testing -r (reverse) execution"},
		{args: []string{"-1sh"}, testFile: "logo-ls-sh.snap", td: "Testing -sh (human readable size) execution"},
		{args: []string{"-1R"}, testFile: "logo-ls-R.snap", td: "Testing -R (recursive) execution"},
		{args: []string{"-1Ra"}, testFile: "logo-ls-Ra.snap", td: "Testing -Ra (recursive, all) execution"},
		{args: []string{"-1shRa"}, testFile: "logo-ls-shRa.snap", td: "Testing -shRa execution"},
		{args: []string{"-V"}, testFile: "logo-ls-V.snap", td: "Testing -V option prints version"},
		{args: []string{"-?"}, testFile: "logo-ls--help.snap", td: "Testing -? (help) prints help message"},
	}

	for _, test := range tt {
		t.Run(test.td, func(st *testing.T) {
			args := []string{"run", "."}
			args = append(args, test.args...)
			cmd := exec.Command(goExec, append(args, dirEnv)...)
			stdout, err := cmd.StdoutPipe()
			defer stdout.Close()
			if err != nil {
				st.Fatal(err)
			}
			if err := cmd.Start(); err != nil {
				st.Fatal(err)
			}
			cmdData, err := ioutil.ReadAll(stdout)
			if err != nil {
				st.Fatal(err)
			}
			if err := cmd.Wait(); err != nil {
				st.Fatal(err)
			}

			fileData, err := ioutil.ReadFile(filepath.Join("./testdata", test.testFile))
			if err != nil {
				st.Fatal(err)
			}

			if bytes.Compare(cmdData, fileData) != 0 {
				t.Fatalf("expected output of the command:\n-----------\n%s\n=============\nbut got:\n-----------\n%s", fileData, cmdData)
			}
		})
	}

}
