package main

import (
	"fmt"
	"os"
	"strings"
)

func mainSort(a, b string) bool {
	switch a {
	case ".", "..":
	default:
		a = strings.TrimPrefix(a, ".")
	}
	switch b {
	case ".", "..":
	default:
		b = strings.TrimPrefix(b, ".")
	}
	return strings.ToLower(a) < strings.ToLower(b)
}

// Custom less functions
func lessFuncGenerator(d *dir) {
	switch {
	case (flagVector & flag_alpha) > 0:
		// sort by alphabetical order of name.ext
		d.less = func(i, j int) bool {
			return mainSort(d.files[i].name+d.files[i].ext, d.files[j].name+d.files[j].ext)
		}
	case (flagVector & flag_S) > 0:
		// sort by file size, largest first
		d.less = func(i, j int) bool {
			if d.files[i].size > d.files[j].size {
				return true
			} else if d.files[i].size == d.files[j].size {
				return mainSort(d.files[i].name+d.files[i].ext, d.files[j].name+d.files[j].ext)
			} else {
				return false
			}
		}
	case (flagVector & flag_t) > 0:
		// sort by modification time, newest first
		// not sorting by alphabetical order because equality is quite rare
		d.less = func(i, j int) bool {
			return d.files[i].modTime.After(d.files[j].modTime)
		}
	case (flagVector & flag_X) > 0:
		// sort alphabetically by entry extension
		d.less = func(i, j int) bool {
			if mainSort(d.files[i].ext, d.files[j].ext) {
				return true
			} else if strings.ToLower(d.files[i].ext) == strings.ToLower(d.files[j].ext) {
				return mainSort(d.files[i].name+d.files[i].ext, d.files[j].name+d.files[j].ext)
			} else {
				return false
			}
		}
	case (flagVector & flag_v) > 0:
		// natural sort of (version) numbers within text
		d.less = func(i, j int) bool {
			return d.files[i].name+d.files[i].ext < d.files[j].name+d.files[j].ext
		}
	default:
		// no sorting
		d.less = func(i, j int) bool {
			return i < j
		}
	}
}

// get Owner and Group info
var grpMap = make(map[string]string)
var userMap = make(map[string]string)

// get indicator of the file
func getIndicator(modebit os.FileMode) (i string) {
	switch {
	case modebit&os.ModeDir > 0:
		i = "/"
	case modebit&os.ModeNamedPipe > 0:
		i = "|"
	case modebit&os.ModeSymlink > 0:
		i = "@"
	case modebit&os.ModeSocket > 0:
		i = "="
	case modebit&1000000 > 0:
		i = "*"
	}
	return i
}

func getSizeInFormate(b int64) string {
	if flagVector&flag_h == 0 {
		return fmt.Sprintf("%d", b)
	}

	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func getIcon(name, ext, indicator string) (icon, color string) {
	var i *iInfo
	var ok bool
	switch indicator {
	case "/":
		// send dir related stuff
		i, ok = Icon_Dir[strings.ToLower(name+ext)]
		if ok {
			break
		}
		if len(name) == 0 || '.' == name[0] {
			i = iDef["hiddendir"]
			break
		}
		i = iDef["dir"]
	case "*":
		// send executable related stuff
		i = iDef["exe"]
	default:
		// send file related stuff
		i, ok = Icon_FileName[strings.ToLower(name+ext)]
		if ok {
			break
		}

		i, ok = Icon_Ext[strings.ToLower(strings.TrimPrefix(ext, "."))]
		if ok {
			break
		}

		if len(name) == 0 || '.' == name[0] {
			i = iDef["hiddenfile"]
			break
		}
		i = iDef["file"]
	}
	return i.getGlyph(), i.getColor(1)
}
