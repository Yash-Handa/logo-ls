package ctw

import "strings"

var (
	noColor string = "\033[0m"
	green   string = "\033[38;2;055;183;021m"
	brown   string = "\033[38;2;192;154;107m"
	empty   string = "\u0020"
)

func DisplayColor(b bool) {
	if b == false {
		noColor = ""
		green = ""
		brown = ""
	}
}

func getGitColor(gitStatus string) string {
	switch strings.Trim(gitStatus, " ") {
	case "":
		return noColor
	case "U":
		return green
	default:
		return brown
	}
}

func widthsSum(w [][4]int, p int) int {
	s := 0
	for _, v := range w {
		s += v[0] + v[1] + v[2] + v[3] + p
	}
	s -= p
	return s
}
