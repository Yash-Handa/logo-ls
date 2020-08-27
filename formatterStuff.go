package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

// Custom less functions
func lessFuncGenerator(d *dir) {
	switch {
	case (flagVector & flag_alpha) > 0:
		// sort by alphabetical order of name.ext
		d.less = func(i, j int) bool {
			return strings.ToLower(d.files[i].name+d.files[i].ext) < strings.ToLower(d.files[j].name+d.files[j].ext)
		}
	case (flagVector & flag_S) > 0:
		// sort by file size, largest first
		d.less = func(i, j int) bool {
			return d.files[i].size > d.files[j].size
		}
	case (flagVector & flag_t) > 0:
		// sort by modification time, newest first
		d.less = func(i, j int) bool {
			return d.files[i].modTime.After(d.files[j].modTime)
		}
	case (flagVector & flag_X) > 0:
		// sort alphabetically by entry extension
		d.less = func(i, j int) bool {
			return strings.ToLower(d.files[i].ext) < strings.ToLower(d.files[j].ext)
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

func getOwnerGroupInfo(fi os.FileInfo) (o string, g string) {
	if stat, ok := fi.Sys().(*syscall.Stat_t); ok {
		if flagVector&(flag_l|flag_o) > 0 {
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

		if flagVector&flag_G == 0 && flagVector&(flag_l|flag_g) > 0 {
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

// get indicator of the file
func getIndicator(modebit os.FileMode) (i string) {
	switch {
	case modebit&os.ModeDir > 0:
		i = "/"
	case modebit&1000000 > 0:
		i = "*"
	case modebit&os.ModeNamedPipe > 0:
		i = "|"
	case modebit&os.ModeSymlink > 0:
		i = "@"
	case modebit&os.ModeSocket > 0:
		i = "="
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
