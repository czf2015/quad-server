package api_v3

import (
	"goserver/libs/orm"
)

var db orm.DB

func init() {
	db = orm.GetDB()
}
