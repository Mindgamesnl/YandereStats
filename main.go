package main

import (
	"github.com/Mindgamesnl/YandereStats/analytics"
	"github.com/Mindgamesnl/YandereStats/changelog"
	"github.com/Mindgamesnl/YandereStats/config"
	"github.com/Mindgamesnl/YandereStats/database"
	"github.com/Mindgamesnl/YandereStats/initializer"
	"github.com/Mindgamesnl/YandereStats/utils"
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
)

var (
	UpdateRepository *git.Repository
	ChangeLog        changelog.ChangeLog
)

func main() {
	total := utils.NewStopwatch("Data Collection - Combined Total")
	config.LoadConfiguration()
	UpdateRepository = initializer.InitializeGit()
	ChangeLog = initializer.InitializeGameVersions()
	ChangeLog = initializer.MergeDataSets(UpdateRepository, ChangeLog)

	database.SaveToSql(ChangeLog)
	total.Stop()

	// analytics
	logrus.Info("Starting analytical tasks")
	bar := pb.StartNew(len(analytics.AnalyticalTasks))
	for i := range analytics.AnalyticalTasks {
		analytics.AnalyticalTasks[i](ChangeLog)
		bar.Increment()
	}
	bar.Finish()

	analytics.GenerateStopwatchBreakdown()
}
