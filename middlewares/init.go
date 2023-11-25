package middlewares

import (
	"goserver/libs/orm"
)

var db orm.DB

func init() {
	db = orm.GetDB()
}
