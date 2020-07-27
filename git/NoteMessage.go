package git

type NoteMessage struct {
	id     uint64 `gorm:"primary_key;auto_increment:true"`
	UpdateID uint
	Message string
}
