package initializer

import (
	"encoding/json"
	"github.com/Mindgamesnl/YandereStats/changelog"
	git2 "github.com/Mindgamesnl/YandereStats/git"
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/format/diff"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"io/ioutil"
	"strconv"
	"strings"
)

func MergeDataSets(repo *git.Repository, cl changelog.ChangeLog) changelog.ChangeLog {
	bar := pb.StartNew(len(cl.Revisions))
	logrus.Info("Merging changelog and repository, this may take a while")

	var failedMatches = 0
	var successfullMatches = 0

	for i := range cl.Revisions {
		rev := &cl.Revisions[i]
		commit := findCommitByName(repo, rev.GitFormattedDate, &cl, rev)
		var valID = len(commit.Changes) > 0
		if !valID {
			failedMatches++
		} else {
			successfullMatches++
			cl.Revisions[i].CommitData = commit
		}
		bar.Increment()
	}

	file, _ := json.MarshalIndent(cl, "", " ")
	_ = ioutil.WriteFile("secrets/merged.json", file, 0644)

	bar.Finish()
	logrus.Info("Finished commit matching. Failed for " + strconv.FormatInt(int64(failedMatches), 10) + " commits and was successful for " + strconv.FormatInt(int64(successfullMatches), 10))
	return cl
}

type convert func(commit *object.Commit)

func findCommitByName(repo *git.Repository, name string, log *changelog.ChangeLog, rev *changelog.ChangelogRevision) git2.Commit {
	ref, _ := repo.Head()
	cIter, _ := repo.Log(&git.LogOptions{From: ref.Hash()})

	var result git2.Commit

	var looper convert

	looper = func(c *object.Commit) {
		if c != nil {
			message := strings.ReplaceAll(c.Message, "\n", "")
			if message == name {

				log.UpdateID++
				ID := log.UpdateID
				out := git2.Commit{UpdateID: ID}
				// set update ID's
				rev.UpdateID = ID
				for i := range rev.Note {
					rev.Note[i].UpdateID = ID
				}

				parentCommit, _ := c.Parents().Next()

				patchSet, _ := c.Patch(parentCommit)
				filePatches := patchSet.FilePatches()
				out.AddedLines = []git2.LineUpdate{}
				out.RemovedLines = []git2.LineUpdate{}

				for i := range filePatches {
					filePatch := filePatches[i]
					chunks := filePatch.Chunks()
					for chunkIterator := range chunks {
						fileChunk := chunks[chunkIterator]
						if fileChunk.Type() == diff.Add {
							out.AddedLines = append(out.AddedLines, git2.LineUpdate{Code: fileChunk.Content(), UpdateID: ID, Action: "ADD"})
						} else if fileChunk.Type() == diff.Delete {
							out.RemovedLines = append(out.RemovedLines, git2.LineUpdate{Code: fileChunk.Content(), UpdateID: ID, Action: "DELETE"})
						}
					}
				}

				stats, _ := c.Stats()
				for i := range stats {
					stat := stats[i]
					file := git2.ChangedFile{FileName: stat.Name, AddedLines: stat.Addition, RemovedLines: stat.Deletion, UpdateId: ID}
					out.Changes = append(out.Changes, file)
				}

				result = out
			} else {
				first, _ := cIter.Next()
				looper(first)
			}
		}
	}

	first, _ := cIter.Next()

	looper(first)
	return result
}
