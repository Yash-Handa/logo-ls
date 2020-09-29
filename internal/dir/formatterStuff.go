package dir

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"

	"github.com/Yash-Handa/logo-ls/assets"
	"github.com/Yash-Handa/logo-ls/internal/api"
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
	case (api.FlagVector & api.Flag_alpha) > 0:
		// sort by alphabetical order of name.ext
		d.less = func(i, j int) bool {
			return mainSort(d.files[i].name+d.files[i].ext, d.files[j].name+d.files[j].ext)
		}
	case (api.FlagVector & api.Flag_S) > 0:
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
	case (api.FlagVector & api.Flag_t) > 0:
		// sort by modification time, newest first
		// not sorting by alphabetical order because equality is quite rare
		d.less = func(i, j int) bool {
			return d.files[i].modTime.After(d.files[j].modTime)
		}
	case (api.FlagVector & api.Flag_X) > 0:
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
	case (api.FlagVector & api.Flag_v) > 0:
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

func getOwnerGroupInfo(fi os.FileInfo) (o string, g string) {
	if stat, ok := fi.Sys().(*syscall.Stat_t); ok {
		if api.FlagVector&(api.Flag_l|api.Flag_o) > 0 {
			UID := strconv.Itoa(int(stat.Uid))
			if n, ok := userMap[UID]; ok {
				o = n
			} else {
				u, err := user.LookupId(UID)
				if err != nil {
					o = ""
				} else {
					o = u.Name
					userMap[UID] = u.Name
				}
			}
		}

		if api.FlagVector&api.Flag_G == 0 && api.FlagVector&(api.Flag_l|api.Flag_g) > 0 {
			GID := strconv.Itoa(int(stat.Gid))
			if n, ok := grpMap[GID]; ok {
				g = n
			} else {
				grp, err := user.LookupGroupId(GID)
				if err != nil {
					g = ""
				} else {
					g = grp.Name
					grpMap[GID] = grp.Name
				}
			}
		}
	}

	return
}

func getFileBlocks(fi os.FileInfo) int64 {
	if s, ok := fi.Sys().(*syscall.Stat_t); ok {
		return s.Blocks
	}
	return 0
}

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
		// send dir related stuff
		i, ok = assets.Icon_Dir[strings.ToLower(name+ext)]
		if ok {
			break
		}
		if len(name) == 0 || '.' == name[0] {
			i = assets.Icon_Def["hiddendir"]
			break
		}
		i = assets.Icon_Def["dir"]
	case "*":
		// send executable related stuff
		i, ok = assets.Icon_FileName[strings.ToLower(name+ext)]
		if ok {
			i.MakeExe()
			break
		}

		i, ok = assets.Icon_Ext[strings.ToLower(strings.TrimPrefix(ext, "."))]
		if ok {
			i.MakeExe()
			break
		}

		if len(name) == 0 || '.' == name[0] {
			i = assets.Icon_Def["hiddenfile"]
			i.MakeExe()
			break
		}
		i = assets.Icon_Def["exe"]
	default:
		// send file related stuff
		i, ok = assets.Icon_FileName[strings.ToLower(name+ext)]
		if ok {
			break
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
	return i.GetGlyph(), i.GetColor(1)
}
