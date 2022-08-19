package models_v2

import "goserver/libs/gorm"

type Publish struct {
	Page
	Version string `json:"version"`
	Remark  string `json:"remark"`
	Online  int    `gorm:"default:0" json:"online"`
}

func init() {
	gorm.AutoMigrat(&Publish{})
}
