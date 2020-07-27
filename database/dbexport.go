package database

import (
	"github.com/Mindgamesnl/YandereStats/changelog"
	"github.com/Mindgamesnl/YandereStats/git"
	"github.com/cheggaaa/pb/v3"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

func SaveToSql(cl changelog.ChangeLog)  {
	bar := pb.StartNew(len(cl.Revisions))
	logrus.Info("Creating database")
	db, err := gorm.Open("sqlite3", "secrets/database.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&git.NoteMessage{})
	db.AutoMigrate(&git.LineUpdate{})
	db.AutoMigrate(&git.ChangedFile{})
	db.AutoMigrate(&changelog.ChangeLog{})
	db.AutoMigrate(&changelog.ChangelogRevision{})
	db.AutoMigrate(&git.Commit{})

	logrus.Info("Saving data to MySQL, this may take a while")
	for i := range cl.Revisions {
		toUpload := cl.Revisions[i]
		if len(toUpload.CommitData.RemovedLines) > 0 {
			db.Create(toUpload)

			// git stuff
			db.Create(toUpload.CommitData)
			for fileChange := range toUpload.CommitData.Changes {
				db.Create(toUpload.CommitData.Changes[fileChange])
			}
			for fileChange := range toUpload.Note {
				db.Create(toUpload.Note[fileChange])
			}
			for fileChange := range toUpload.CommitData.AddedLines {
				db.Create(toUpload.CommitData.AddedLines[fileChange])
			}
			for fileChange := range toUpload.CommitData.RemovedLines {
				db.Create(toUpload.CommitData.RemovedLines[fileChange])
			}

			if db.Error != nil {
				logrus.Error(db.Error)
			}
		}
		bar.Increment()
	}
	bar.Finish()
}
