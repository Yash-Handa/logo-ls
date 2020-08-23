package main

import (
	"os"

	"github.com/pborman/getopt/v2"
)

func main() {
	helpFlag := getopt.Bool('?', "display this help and exit")
	_ = getopt.BoolLong("all", 'a', "do not ignore entries starting with .")
	_ = getopt.BoolLong("almost-all", 'A', "do not list implied . and ..")
	_ = getopt.Bool('1', "list one file per line.  Avoid '\\n' with -q or -b")
	_ = getopt.BoolLong("version", 'V', "output version information and exit")
	getopt.ParseV2()
	if *helpFlag {
		getopt.PrintUsage(os.Stdout)
	}
}
