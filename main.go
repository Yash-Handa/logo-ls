package main

import (
	"os"

	"github.com/pborman/getopt/v2"
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

	helpFlag := getopt.Bool('?', "display this help and exit")
	_ = getopt.BoolLong("version", 'V', "output version information and exit")

	getopt.ParseV2()
	if *helpFlag {
		getopt.PrintUsage(os.Stdout)

	}
}
