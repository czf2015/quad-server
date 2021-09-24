package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"goserver/libs/conf"
)

var db *gorm.DB

func init() {
	dbCfg, err := conf.GetSection("database")

	dbType := dbCfg.Key("TYPE").String()
	dbName := dbCfg.Key("NAME").String()
	user := dbCfg.Key("USER").String()
	password := dbCfg.Key("PASSWORD").String()
	host := dbCfg.Key("HOST").String()
	charset := dbCfg.Key("CHARSET").String()
	parseTime := dbCfg.Key("PARSE_TIME").String()
	loc := dbCfg.Key("LOC").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		user,
		password,
		host,
		dbName,
		charset,
		parseTime,
		loc,
	))
	if err != nil {
		log.Println(err)
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func Create(value interface{}) *gorm.DB {
	return db.Create(value)
}

func DB() *gorm.DB {
	return db
}