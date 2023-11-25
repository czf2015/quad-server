package models_v3

type Publish struct {
	Page
	Version string `gorm:"type:varchar(255);not null" json:"version"`
	Remark  string `gorm:"type:varchar(255)" json:"remark"`
	Online  int    `gorm:"default:0" json:"online"`
}
