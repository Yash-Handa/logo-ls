// this file contain dir type definition

// +build !windows

package main

import (
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func dirBlocks(info *file, fi os.FileInfo) {
	if s, ok := fi.Sys().(*syscall.Stat_t); ok {
		info.blocks = s.Blocks
	}
}

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
