package dir

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

// gitStatus returns the status of modified files in the worktree. It will attempt to execute 'git status'
// and will fall back to git.Worktree.Status() if that fails.
func gitStatus(wt *git.Worktree) (git.Status, error) {
	c := exec.Command("git", "status", "--porcelain", "-z")
	c.Dir = wt.Filesystem.Root()
	output, err := c.Output()
	if err != nil {
		stat, err := wt.Status()
		return stat, err
	}

	lines := strings.Split(string(output), "\000")
	stat := make(map[string]*git.FileStatus, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		// For copy/rename the output looks like
		//   R  destination\000source
		// Which means we can split on space and ignore anything with only one result
		parts := strings.SplitN(strings.TrimLeft(line, " "), " ", 2)
		if len(parts) == 2 {
			stat[strings.Trim(parts[1], " ")] = &git.FileStatus{
				Worktree: git.StatusCode([]byte(parts[0])[0]),
			}
		}
	}
	return stat, err
}

var (
	gitRepo                git.Status
	gitRoot                string
	gitError               error
	gitRepoComputedAlready = false
)

func getRepoStatus(path string) (git.Status, string, error) {
	if !gitRepoComputedAlready {
		gitRepoComputedAlready = true
		op := git.PlainOpenOptions{DetectDotGit: true}
		r, err := git.PlainOpenWithOptions(path, &op)
		if err != nil {
			gitError = err
			return nil, "", err
		}

		w, err := r.Worktree()
		if err != nil {
			gitError = err
			return nil, "", err
		}

		ws, err := gitStatus(w)
		if err != nil {
			gitError = err
			return nil, "", err
		}

		gitRepo, gitRoot, gitError = ws, w.Filesystem.Root(), nil
	}
	return gitRepo, gitRoot, gitError
}

// use this function when multiple dirs are given as input, to recalculate the git tree for each
func GitRepoCompute() {
	gitRepoComputedAlready = false
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
