// This file has sysState data that is set and used by other file
package sysState

// terminal width for formatting [default to 80]
var terminalWidth int = 80

func TerminalWidth(w int) {
	terminalWidth = w
}

func GetTerminalWidth() int {
	return terminalWidth
}

const (
	Code_OK int = iota
	Code_Minor
	Code_Serious
)

// os exit code (do not update manually)
var osExitCode int = Code_OK

// only use set_osExitCode to update the value of osExitCode
func ExitCode(c int) {
	switch {
	case c == Code_Serious:
		osExitCode = Code_Serious
	case c == Code_Minor && osExitCode != Code_Serious:
		osExitCode = Code_Minor
	}

}

func GetExitCode() int {
	return osExitCode
}
