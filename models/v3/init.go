package models_v3

import (
	"goserver/libs/conf"
	"goserver/libs/orm"
)

var appUrl string
var db orm.DB

func init() {
	appUrl = conf.GetSectionKey("app", "APP_URL").String()

	db = orm.GetDB()
	db.AutoMigrate(&User{}, &Role{}, &Permission{}, &Page{}, &Template{}, &Publish{}, &Menu{}, &EventRule{})
}
