package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pborman/getopt/v2"
)

// flags with corresponding bit values
// frequently used flags should be higher in the list
// help (-?) and version (-V) not included
const (
	flag_a uint = 1 << iota
	flag_l
	flag_alpha // sort in alphabetic order (default)
	flag_A
	flag_h
	flag_1
	flag_d
	flag_s
	flag_r
	flag_S
	flag_t
	flag_X
	flag_v
	flag_U
	flag_o
	flag_g
	flag_G
)

func main() {
	// content flags
	_ = getopt.BoolLong("all", 'a', "do not ignore entries starting with .")
	_ = getopt.BoolLong("almost-all", 'A', "do not list implied . and ..")

	// display flags
	_ = getopt.Bool('1', "list one file per line.")
	_ = getopt.BoolLong("directory", 'd', "list directories themselves, not their contents")
	_ = getopt.Bool('l', "use a long listing format")
	_ = getopt.Bool('o', "like -l, but do not list group information")
	_ = getopt.Bool('g', "\nlike -l, but do not list owner")
	_ = getopt.BoolLong("no-group", 'G', "in a long listing, don't print group names")
	_ = getopt.BoolLong("human-readable", 'h', "with -l and -s, print sizes like 1K 234M 2G etc.")
	_ = getopt.BoolLong("size", 's', "print the allocated size of each file, in blocks") //use os.Getpagesize()

	// sorting flags
	_ = getopt.Bool('S', "sort by file size, largest first")
	_ = getopt.Bool('U', "do not sort; list entries in directory order")
	_ = getopt.Bool('X', "sort alphabetically by entry extension")
	_ = getopt.Bool('v', "natural sort of (version) numbers within text")
	_ = getopt.Bool('t', "sort by modification time, newest first")

	_ = getopt.BoolLong("reverse", 'r', "reverse order while sorting")

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
		fmt.Printf("logo-ls %s\nCopyright (c) 2020 Yash Handa\nLicense MIT <https://opensource.org/licenses/MIT>.\nThis is free software: you are free to change and redistribute it.\nThere is NO WARRANTY, to the extent permitted by law.\n", "v1.0.0")
		os.Exit(0)
	}
}

func init() {
	getopt.SetParameters("[files ...]")
	log.SetPrefix("logo-ls: ")
	log.SetFlags(0)
}
