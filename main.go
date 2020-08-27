package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pborman/getopt/v2"
)

// flags with corresponding bit values
// frequently used flags should be higher in the list
// help (-?) and version (-V) not included
const (
	flag_l uint = 1 << iota
	flag_a
	flag_alpha // sort in alphabetic order (default)
	flag_A
	flag_h
	flag_r
	flag_S
	flag_t
	flag_X
	flag_s
	flag_v
	flag_U
	flag_1
	flag_d
	flag_o
	flag_g
	flag_G
)

// flagVector has all the options set in it. Each bit represent an option.
var flagVector uint

func main() {
	// content flags
	f_a := getopt.BoolLong("all", 'a', "do not ignore entries starting with .")
	f_A := getopt.BoolLong("almost-all", 'A', "do not list implied . and ..")

	// display flags
	f_1 := getopt.Bool('1', "list one file per line.")
	f_d := getopt.BoolLong("directory", 'd', "list directories themselves, not their contents")
	f_l := getopt.Bool('l', "use a long listing format")
	f_o := getopt.Bool('o', "like -l, but do not list group information")
	f_g := getopt.Bool('g', "\nlike -l, but do not list owner")
	f_G := getopt.BoolLong("no-group", 'G', "in a long listing, don't print group names")
	f_h := getopt.BoolLong("human-readable", 'h', "with -l and -s, print sizes like 1K 234M 2G etc.")
	f_s := getopt.BoolLong("size", 's', "print the allocated size of each file, in blocks")

	// sorting flags
	f_S := getopt.Bool('S', "sort by file size, largest first")
	f_U := getopt.Bool('U', "do not sort; list entries in directory order")
	f_X := getopt.Bool('X', "sort alphabetically by entry extension")
	f_v := getopt.Bool('v', "natural sort of (version) numbers within text")
	f_t := getopt.Bool('t', "sort by modification time, newest first")

	f_r := getopt.BoolLong("reverse", 'r', "reverse order while sorting")

	f_help := getopt.Bool('?', "display this help and exit")
	f_V := getopt.BoolLong("version", 'V', "output version information and exit")

	// using getopt.Getopt instead of parse to provide custom err
	err := getopt.Getopt(nil)
	if err != nil {
		// code to handle error
		log.Printf("%v\nTry 'logo-ls -?' for more information.", err)
		os.Exit(2)
	}

	// if f_help is provided print help and exit(0)
	if *f_help {
		getopt.PrintUsage(os.Stdout)
		os.Exit(0)
	}

	// if f_V is provided version will be printed and exit(0)
	if *f_V {
		fmt.Printf("logo-ls %s\nCopyright (c) 2020 Yash Handa\nLicense MIT <https://opensource.org/licenses/MIT>.\nThis is free software: you are free to change and redistribute it.\nThere is NO WARRANTY, to the extent permitted by law.\n", "v0.0.0")
		os.Exit(0)
	}

	// set one of -A and -a priority -A > -a
	switch {
	case *f_A:
		flagVector |= flag_A
	case *f_a:
		flagVector |= flag_a
	}

	// set one of -S, -U, -X, -v, -t and alpha priority -S > -t > -X > -v > -U > alpha
	switch {
	case *f_S:
		flagVector |= flag_S
	case *f_t:
		flagVector |= flag_t
	case *f_X:
		flagVector |= flag_X
	case *f_v:
		flagVector |= flag_v
	case *f_U:
		flagVector |= flag_U
	default:
		flagVector |= flag_alpha
	}

	// set reverse (-r) flag
	if *f_r {
		flagVector |= flag_r
	}

	// set -1 flag
	if *f_1 {
		flagVector |= flag_1
	}

	// set -d flag
	if *f_d {
		flagVector |= flag_d
	}

	// set -G flag
	if *f_G {
		flagVector |= flag_G
	}

	// set -h flag
	if *f_h {
		flagVector |= flag_h
	}

	// set -s flag
	if *f_s {
		flagVector |= flag_s
	}

	// set one of -o, -g and -l priority -o > -g > -l
	switch {
	case *f_o:
		flagVector |= flag_o
	case *f_g:
		flagVector |= flag_g
	case *f_l:
		flagVector |= flag_l
	}

	// extract files/dir from arguments
	dirs := getopt.Args()
	if len(dirs) == 0 {
		// use pwd
		pwd, err := os.Open(".")
		if err != nil {
			log.Printf("cannot access \".\": %v\n", err)
			pwd.Close()
			os.Exit(2)
		}
		d, err := newDir(pwd)
		pwd.Close()
		if err != nil {
			log.Printf("partial access to \".\": %v\n", err)
			defer os.Exit(2)
		}
		// print the info of the files of pwd
		io.Copy(os.Stdout, d.print())
	} else {
		for _, v := range dirs {
			// use this list containing both dirs and files
			_ = v
		}
	}
}

func init() {
	getopt.SetParameters("[files ...]")
	log.SetPrefix("logo-ls: ")
	log.SetFlags(0)
}

// todo: i. multiple dir/ file handle
