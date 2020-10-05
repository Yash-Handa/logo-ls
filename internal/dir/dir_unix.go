// this file contain dir type definition

// +build !windows

package dir

import (
	"os"
	"os/user"
	"strconv"
	"syscall"

	"github.com/Yash-Handa/logo-ls/internal/api"
)

func dirBlocks(info *file, fi os.FileInfo) {
	if s, ok := fi.Sys().(*syscall.Stat_t); ok {
		info.blocks = s.Blocks
	}
}

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
