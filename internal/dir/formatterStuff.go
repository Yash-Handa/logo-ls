package dir

import (
	"fmt"
	"os"
	"strings"

	"github.com/Yash-Handa/logo-ls/assets"
	"github.com/Yash-Handa/logo-ls/internal/api"
	"golang.org/x/exp/constraints"
)

func threeWayCompare[K constraints.Ordered](i,j K) int {
	if i < j {
		return -1
	}
	if j < i {
		return 1
	}
	return 0
}

func cleanupName(name string) string {
	switch name {
	case ".", "..":
	default:
		name = strings.TrimPrefix(name, ".")
	}

	return strings.ToLower(name)
}

func indexCmp(d *dir, i, j int) int {
	return threeWayCompare(i, j)
}

func nameCmp(d *dir, i, j int) int {
	return strings.Compare(cleanupName(d.files[i].name + d.files[i].ext), cleanupName(d.files[j].name + d.files[j].ext))
}

func naturalCmp(d *dir, i, j int) int {
	return threeWayCompare(d.files[i].name + d.files[i].ext, d.files[j].name + d.files[j].ext)
}

func extCmp(d *dir, i, j int) int {
	return strings.Compare(cleanupName(d.files[i].ext), cleanupName(d.files[j].ext))
}

func sizeCmp(d *dir, i, j int) int {
	// size comparison is largest first, so negate the result
	return -threeWayCompare(d.files[i].size, d.files[j].size)
}

func dirFlagCmp(d *dir, i, j int) int {
	// directory flag compare, directories first
	if d.files[i].isDir == d.files[j].isDir {
		return 0
	}
	if d.files[i].isDir {
		return -1
	}
	return 1
}

func dirFlagCmpReverse(d *dir, i, j int) int {
	return -dirFlagCmp(d, i, j)
}

func modTimeCmp(d *dir, i, j int) int {
	if d.files[i].modTime.After(d.files[j].modTime) {
		return -1
	}
	if d.files[i].modTime.Before(d.files[j].modTime) {
		return 1
	}
	return 0
}

func doCompare(d *dir, i, j int, funcs [] func(d *dir, i,j int) int) bool {
	for k := len(funcs)-1; k >= 0; k-- {
		var result = funcs[k](d, i, j)
		switch result {
		case -1: return true
		case 1: return false
		case 0: // do nothing, go to the next comparator
		}
	}
	return false
}


// Custom less functions
func lessFuncGenerator(d *dir) {
	var compareFuncs = []func(d *dir, i,j int) int{}

	switch {
	case (api.FlagVector & api.Flag_alpha) > 0:
		compareFuncs = append(compareFuncs, nameCmp)
	case (api.FlagVector & api.Flag_S) > 0:
		// sort by file size, largest first
		compareFuncs = append(compareFuncs, nameCmp, sizeCmp)
	case (api.FlagVector & api.Flag_t) > 0:
		// sort by modification time, newest first
		// not sorting by alphabetical order because equality is quite rare
		compareFuncs = append(compareFuncs, modTimeCmp)
	case (api.FlagVector & api.Flag_X) > 0:
		// sort alphabetically by entry extension
		compareFuncs = append(compareFuncs, nameCmp, extCmp)
	case (api.FlagVector & api.Flag_v) > 0:
		// natural sort of (version) numbers within text
		compareFuncs = append(compareFuncs, naturalCmp)
	default:
		compareFuncs = append(compareFuncs, indexCmp)
	}

	if (api.FlagVector & api.Flag_groupDirs) > 0 {
		if (api.FlagVector & api.Flag_r) > 0 {
			// Reverse the order of dir vs files if reverse flag is on
			compareFuncs = append(compareFuncs, dirFlagCmpReverse)
		} else {
			compareFuncs = append(compareFuncs, dirFlagCmp)
		}
	}

	d.less = func(i, j int) bool {
		return doCompare(d, i, j, compareFuncs)
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
	if api.FlagVector&api.Flag_h == 0 {
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
	var i *assets.Icon_Info
	var ok bool

	switch indicator {
	case "/":
		i, ok = assets.Icon_Dir[strings.ToLower(name+ext)]
		if ok {
			break
		}
		if len(name) == 0 || '.' == name[0] {
			i = assets.Icon_Def["hiddendir"]
			break
		}
		i = assets.Icon_Def["dir"]
	default:
		i, ok = assets.Icon_FileName[strings.ToLower(name+ext)]
		if ok {
			break
		}

		// a special admiration for goLang
		if ext == ".go" && strings.HasSuffix(name, "_test") {
			i = assets.Icon_Set["go-test"]
			break
		}

		t := strings.Split(name, ".")
		if len(t) > 1 && t[0] != "" {
			i, ok = assets.Icon_SubExt[strings.ToLower(t[len(t)-1]+ext)]
			if ok {
				break
			}
		}

		i, ok = assets.Icon_Ext[strings.ToLower(strings.TrimPrefix(ext, "."))]
		if ok {
			break
		}

		if len(name) == 0 || '.' == name[0] {
			i = assets.Icon_Def["hiddenfile"]
			break
		}
		i = assets.Icon_Def["file"]
	}

	// change icon color if the file is executable
	if indicator == "*" {
		if i.GetGlyph() == "\uf723" {
			i = assets.Icon_Def["exe"]
		}
		i.MakeExe()
	}

	return i.GetGlyph(), i.GetColor(1)
}
