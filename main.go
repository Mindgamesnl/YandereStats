package main

import (
	"github.com/Mindgamesnl/YandereFetch/changelog"
	"github.com/Mindgamesnl/YandereFetch/config"
	"github.com/Mindgamesnl/YandereFetch/database"
	"github.com/Mindgamesnl/YandereFetch/graphing"
	"github.com/Mindgamesnl/YandereFetch/initializer"
	"gopkg.in/src-d/go-git.v4"
)

var UpdateRepository *git.Repository
var ChangeLog changelog.ChangeLog

func main() {
	config.LoadConfiguration()
	UpdateRepository = initializer.InitializeGit()
	ChangeLog = initializer.InitializeGameVersions()
	ChangeLog = initializer.MergeDataSets(UpdateRepository, ChangeLog)

	database.SaveToSql(ChangeLog)
	graphing.GenerateGraph(ChangeLog)
}