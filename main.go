package main

import (
	"github.com/Mindgamesnl/YandereStats/analytics"
	"github.com/Mindgamesnl/YandereStats/changelog"
	"github.com/Mindgamesnl/YandereStats/config"
	"github.com/Mindgamesnl/YandereStats/initializer"
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
)

var UpdateRepository *git.Repository
var ChangeLog changelog.ChangeLog

func main() {
	config.LoadConfiguration()
	UpdateRepository = initializer.InitializeGit()
	ChangeLog = initializer.InitializeGameVersions()
	ChangeLog = initializer.MergeDataSets(UpdateRepository, ChangeLog)

	// database.SaveToSql(ChangeLog)

	// analytics
	logrus.Info("Starting analytical tasks")
	bar := pb.StartNew(len(analytics.AnalyticalTasks))
	for i := range analytics.AnalyticalTasks {
		analytics.AnalyticalTasks[i](ChangeLog)
		bar.Increment()
	}
	bar.Finish()
}
