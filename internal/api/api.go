//This file contains the cli API and configs it
package api

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Yash-Handa/logo-ls/internal/sysState"
	"github.com/pborman/getopt/v2"
	"golang.org/x/crypto/ssh/terminal"
)

// flags with corresponding bit values
// frequently used flags should be higher in the list
// help (-?) and version (-V) not included
const (
	Flag_l uint = 1 << iota
	Flag_a
	Flag_alpha // sort in alphabetic order (default)
	Flag_i     // stop printing of icons
	Flag_c     // stop printing of colors
	Flag_D     // stop printing of git status
	Flag_A
	Flag_h
	Flag_R
	Flag_r
	Flag_S
	Flag_t
	Flag_X
	Flag_s
	Flag_v
	Flag_U
	Flag_1
	Flag_d
	Flag_o
	Flag_g
	Flag_G
)

// flagVector has all the options set in it. Each bit represent an option.
var FlagVector uint

// time formate
var timeFormate string

func TimeFormate(t string) {
	timeFormate = t
}

func GetTimeFormate() string {
	return timeFormate
}

var FileList []string

func Bootstrap() {
	getopt.SetParameters("[files ...]")

	// content flags
	f_a := getopt.BoolLong("all", 'a', "do not ignore entries starting with .")
	f_A := getopt.BoolLong("almost-all", 'A', "do not list implied . and ..")

	// disable Stuff
	f_D := getopt.BoolLong("git-status", 'D', "print git status of files")
	f_c := getopt.BoolLong("disable-color", 'c', "don't color icons, filenames and git status (use this to print to a file)")
	f_i := getopt.BoolLong("disable-icon", 'i', "don't print icons of the files")

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
	f_R := getopt.BoolLong("recursive", 'R', "list subdirectories recursively")
	f_T := getopt.EnumLong("time-style", 'T', []string{"Stamp", "StampMilli", "Kitchen", "ANSIC", "UnixDate", "RubyDate", "RFC1123", "RFC1123Z", "RFC3339", "RFC822", "RFC822Z", "RFC850"}, "Stamp", "time/date format with -l; see time-style below")

	f_help := getopt.BoolLong("help", '?', "display this help and exit")
	f_V := getopt.BoolLong("version", 'V', "output version information and exit")

	// using getopt.Getopt instead of parse to provide custom err
	err := getopt.Getopt(nil)
	if err != nil {
		// code to handle error
		log.Printf("%v\nTry 'logo-ls -?' for more information.", err)
		sysState.ExitCode(sysState.Code_Serious)
		os.Exit(sysState.GetExitCode())
	}

	// list of files/ dir
	FileList = getopt.Args()

	// set one of -A and -a priority -A > -a
	switch {
	case *f_A:
		FlagVector |= Flag_A
	case *f_a:
		FlagVector |= Flag_a
	}

	// set one of -S, -U, -X, -v, -t and alpha priority -S > -t > -X > -v > -U > alpha
	switch {
	case *f_S:
		FlagVector |= Flag_S
	case *f_t:
		FlagVector |= Flag_t
	case *f_X:
		FlagVector |= Flag_X
	case *f_v:
		FlagVector |= Flag_v
	case *f_U:
		FlagVector |= Flag_U
	default:
		FlagVector |= Flag_alpha
	}

	// set reverse (-r) flag
	if *f_r {
		FlagVector |= Flag_r
	}

	// set recursion (-R) flag
	if *f_R {
		FlagVector |= Flag_R
	}

	// set disable-git-status (-D) flag
	if *f_D {
		FlagVector |= Flag_D
	}

	// set disable-color (-c) flag
	if *f_c {
		FlagVector |= Flag_c
	}

	// set disable-icon (-i) flag
	if *f_i {
		FlagVector |= Flag_i
	}

	// set -1 flag
	if *f_1 {
		FlagVector |= Flag_1
	}

	// set -d flag
	if *f_d {
		FlagVector |= Flag_d
	}

	// set -G flag
	if *f_G {
		FlagVector |= Flag_G
	}

	// set time formate
	switch *f_T {
	case "Stamp":
		timeFormate = time.Stamp
	case "StampMilli":
		timeFormate = time.StampMilli
	case "Kitchen":
		timeFormate = time.Kitchen
	case "ANSIC":
		timeFormate = time.ANSIC
	case "UnixDate":
		timeFormate = time.UnixDate
	case "RubyDate":
		timeFormate = time.RubyDate
	case "RFC1123":
		timeFormate = time.RFC1123
	case "RFC1123Z":
		timeFormate = time.RFC1123Z
	case "RFC3339":
		timeFormate = time.RFC3339
	case "RFC822":
		timeFormate = time.RFC822
	case "RFC822Z":
		timeFormate = time.RFC822Z
	case "RFC850":
		timeFormate = time.RFC850
	default:
		timeFormate = time.Stamp
	}

	// set -h flag
	if *f_h {
		FlagVector |= Flag_h
	}

	// set -s flag
	if *f_s {
		FlagVector |= Flag_s
	}

	// set one of -o, -g and -l priority -o > -g > -l
	switch {
	case *f_o:
		FlagVector |= Flag_o
	case *f_g:
		FlagVector |= Flag_g
	case *f_l:
		FlagVector |= Flag_l
	case *f_1:
	default:
		// screen width for custom tw
		w, _, e := terminal.GetSize(int(os.Stdout.Fd()))
		if e == nil {
			if w == 0 {
				// for systems that don’t support ‘TIOCGWINSZ’.
				w, _ = strconv.Atoi(os.Getenv("COLUMNS"))
			}
			sysState.TerminalWidth(w)
		}
	}

	// if f_help is provided print help and exit(0)
	if *f_help {
		fmt.Println("List information about the FILEs with ICONS and GIT STATUS (the current dir \nby default). Sort entries alphabetically if none of -tvSUX is specified.")

		getopt.PrintUsage(os.Stdout)

		fmt.Println("\nPossible value for --time-style (-T)")
		fmt.Printf(" %-11s %-32q\n", "ANSIC", "Mon Jan _2 15:04:05 2006")
		fmt.Printf(" %-11s %-32q\n", "UnixDate", "Mon Jan _2 15:04:05 MST 2006")
		fmt.Printf(" %-11s %-32q\n", "RubyDate", "Mon Jan 02 15:04:05 -0700 2006")
		fmt.Printf(" %-11s %-32q\n", "RFC822", "02 Jan 06 15:04 MST")
		fmt.Printf(" %-11s %-32q\n", "RFC822Z", "02 Jan 06 15:04 -0700")
		fmt.Printf(" %-11s %-32q\n", "RFC850", "Monday, 02-Jan-06 15:04:05 MST")
		fmt.Printf(" %-11s %-32q\n", "RFC1123", "Mon, 02 Jan 2006 15:04:05 MST")
		fmt.Printf(" %-11s %-32q\n", "RFC1123Z", "Mon, 02 Jan 2006 15:04:05 -0700")
		fmt.Printf(" %-11s %-32q\n", "RFC3339", "2006-01-02T15:04:05Z07:00")
		fmt.Printf(" %-11s %-32q\n", "Kitchen", "3:04PM")
		fmt.Printf(" %-11s %-32q [Default]\n", "Stamp", "Mon Jan _2 15:04:05")
		fmt.Printf(" %-11s %-32q\n", "StampMilli", "Jan _2 15:04:05.000")

		fmt.Println("\nExit status:")
		fmt.Println(" 0  if OK,")
		fmt.Println(" 1  if minor problems (e.g., cannot access subdirectory),")
		fmt.Println(" 2  if serious trouble (e.g., cannot access command-line argument).")
		os.Exit(sysState.GetExitCode())
	}

	// if f_V is provided version will be printed and exit(0)
	if *f_V {
		fmt.Printf("logo-ls %s\nCopyright (c) 2020 Yash Handa\nLicense MIT <https://opensource.org/licenses/MIT>.\nThis is free software: you are free to change and redistribute it.\nThere is NO WARRANTY, to the extent permitted by law.\n", "v1.3.7")
		fmt.Println("\nWritten by Yash Handa")
		os.Exit(sysState.GetExitCode())
	}
}
