package gorm

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

	// DSN格式：[username[:password]@][protocol[(address)]]/gormname[?param1=value1&...&paramN=valueN]
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s", user, password, host, dbName, charset, parseTime, loc)
	db, err = gorm.Open(dbType, dsn)
	if err != nil {
		log.Println(err)
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	defer db.Close()
}

// 执行sql原语
func Exec(result interface {}, sql string, args...interface{}) *gorm.DB {
	return db.Raw(sql, args).Scan(result)
}

// 插入
func Create(record interface{}) *gorm.DB {
	return db.Create(record)
}

// 查询
func Take(result interface{}) *gorm.DB {
	return db.Take(result)
}

func First(result interface{}) *gorm.DB {
	return db.First(result)
}

func Last(result interface{}) *gorm.DB {
	return db.Last(result)
}

func Find(results interface{}) *gorm.DB {
	return db.Find(results)
}

func Pluck(model interface{}, field string, results *[]interface{}) *gorm.DB {
	return db.Model(model).Pluck(field, results)
}

func Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.Where(query, args)
}

func Select(model interface{}, statement string) *gorm.DB {
	return db.Model(model).Select(statement)
}

func Count(model interface{}, total *int) *gorm.DB {
	return db.Model(model).Count(&total)
}

func Cursor(result interface{}, limit, offset int) *gorm.DB {
	return db.Order("create_time desc").Limit(limit).Offset(offset).Find(result)
}

// 更新 
// 1. 保存模型变量值
func Save(record interface{}) *gorm.DB {
	return db.Save(record)
}
// 2. 更新单个字段值
func Update(model interface{}, field string, value interface{}) *gorm.DB {
	return db.Model(model).Update(field, value)
}
// 3. 更新多个字段值
func Updates(model interface{}, updates interface{}) *gorm.DB {
	return db.Model(model).Updates(updates)
}

// 删除
// 1. 用法：db.Where(条件表达式).Delete(空模型变量指针)
func Delete(model interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return db.Where(query, args).Delete(model)
}


