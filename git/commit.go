package git

type Commit struct {
	id           uint64 `gorm:"primary_key;auto_increment:true"`
	UpdateID     uint
	Changes      []ChangedFile
	AddedLines   []LineUpdate
	RemovedLines []LineUpdate
}

type ChangedFile struct {
	id           uint64 `gorm:"primary_key;auto_increment:true"`
	UpdateId     uint
	FileName     string
	AddedLines   int
	RemovedLines int
}
