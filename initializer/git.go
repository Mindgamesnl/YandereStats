package initializer

import (
	"github.com/Mindgamesnl/YandereStats/config"
	"github.com/Mindgamesnl/YandereStats/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
	"strconv"
)

const repoDir = "./secrets/yandere-git/"

func InitializeGit() *git.Repository {
	gitTimer := utils.NewStopwatch("Decompiled History - Combined Total")
	repoTimer := utils.NewStopwatch("Decompiled History - Opening Repo")
	var r *git.Repository

	if exists(repoDir) {
		logrus.Info("Found repo, opening and doing pull")
		opened, _ := git.PlainOpen(repoDir)
		worktree, _ := opened.Worktree()
		_ = worktree.Pull(&git.PullOptions{RemoteName: "origin"})
		r = opened
	} else {
		logrus.Info("Repo not found, trying to clone")
		cloned, _ := git.PlainClone(repoDir, false, &git.CloneOptions{
			URL: config.LoadedInstance.HistoricalDataSource,
			Progress: os.Stdout,
		})

		r = cloned
	}

	repoTimer.Stop()
	headTimer := utils.NewStopwatch("Decompiled History - Fetching HEAD")

	ref, _ := r.Head()
	cIter, _ := r.Log(&git.LogOptions{From: ref.Hash()})
	count := 0
	_ = cIter.ForEach(func(c *object.Commit) error {
		count++
		return nil
	})

	logrus.Info("ValIDated repo, contains " + strconv.Itoa(count) + " updates.")
	headTimer.Stop()
	gitTimer.Stop()
	return r
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil { return true }
	if os.IsNotExist(err) { return false }
	return false
}
