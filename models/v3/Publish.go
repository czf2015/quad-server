package models_v3

import (
	"goserver/libs/orm"
)

type Publish struct {
	Page
	Version string `json:"version"`
	Remark  string `json:"remark"`
	Online  int    `gorm:"default:0" json:"online"`
}

func init() {
	orm.GetDB().AutoMigrate(&Publish{})
}
