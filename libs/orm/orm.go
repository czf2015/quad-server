package orm

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"goserver/libs/utils"
)

type DB = *gorm.DB

var db DB
var sqlDB *sql.DB

type DSN struct {
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Name      string `yaml:"dbName"`
	Charset   string `yaml:"charset"`
	ParseTime string `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
	Port      string `yaml:"port"`
}

func (dsn DSN) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", dsn.Username, dsn.Password, dsn.Host, dsn.Port, dsn.Name, dsn.Charset, dsn.ParseTime, dsn.Loc)
}

func init() {
	var (
		err error
		dsn DSN
	)
	utils.YAML.Unmarshal(utils.ReadFile("conf/database.yml"), &dsn)
	fmt.Println("dsn.String()", dsn.String())
	db, err = gorm.Open(mysql.Open(dsn.String()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{

			// TablePrefix: "t_",   // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			// NameReplacer: strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称。
		},
	})
	if err != nil {
		log.Fatalf("Open database fail: %v", err)
		return
	}

	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	}
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	defer sqlDB.Close()
}
