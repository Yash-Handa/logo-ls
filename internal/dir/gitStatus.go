package dir

import (
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func getRepoStatus(path string) (git.Status, string, error) {
	op := git.PlainOpenOptions{DetectDotGit: true}
	r, err := git.PlainOpenWithOptions(path, &op)
	if err != nil {
		return nil, "", err
	}

	w, err := r.Worktree()
	if err != nil {
		return nil, "", err
	}

	ws, err := w.Status()
	if err != nil {
		return nil, "", err
	}

	return ws, w.Filesystem.Root(), nil
}

func getFilesGitStatus(p string) map[string]string {
	gitRepo, gitRoot, err := getRepoStatus(p)
	if err != nil || len(gitRepo) == 0 {
		return nil
	}
	pAbs, err := filepath.Abs(p)
	if err != nil {
		return nil
	}

	t := make(map[string]string)

	for i, v := range gitRepo {
		i = gitFilePath(gitRoot+"/"+i, pAbs+"/")
		if i == "" {
			continue
		}
		dirs := strings.SplitAfter(i, "/")
		d := ""
		for j, seg := range dirs {
			if j == len(dirs)-1 {
				if v.Worktree == '?' {
					t[i] = "U"
				} else {
					t[i] = string(v.Worktree)
				}
			} else {
				d += seg
				t[d] = "●"
			}
		}
	}
	return t
}

func gitFilePath(gitpath, dirpath string) string {
	if strings.HasPrefix(gitpath, dirpath) {
		return strings.TrimPrefix(gitpath, dirpath)
	}
	return ""
}
