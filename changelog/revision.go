package changelog

import (
	git2 "github.com/Mindgamesnl/YandereStats/git"
	"time"
)

type ChangelogRevision struct {
	id     uint64 `gorm:"primary_key;auto_increment:true"`
	UpdateID           uint
	Date             time.Time
	FormattedWebDate string
	GitFormattedDate string
	Note             []git2.NoteMessage
	CommitData       git2.Commit
}

type ChangeLog struct {
	id     uint64 `gorm:"primary_key;auto_increment:true"`
	UpdateID    uint
	Revisions []ChangelogRevision
}
