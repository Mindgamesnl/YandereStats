package git

type LineUpdate struct {
	id     uint64 `gorm:"primary_key;auto_increment:true"`
	UpdateID uint
	Code string
	Action string
}
