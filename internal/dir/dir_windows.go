// this file contain dir type definition

// +build windows

package dir

import (
	"os"
)

func dirBlocks(info *file, fi os.FileInfo) {
}

func getOwnerGroupInfo(fi os.FileInfo) (o string, g string) {
	return
}
