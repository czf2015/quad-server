package gorm

import (
	"fmt"
	"log"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/driver/mysql"

	"goserver/libs/conf"
)

type DB = gorm.DB

type Model = gorm.Model

var db *gorm.DB
var sqlDB *sql.DB

// DSN格式：[username[:password]@][protocol[(address)]]/gormname[?param1=value1&...&paramN=valueN]
func getDSN() string {
	dbCfg, _ := conf.GetSection("database")

	dbName := dbCfg.Key("NAME").String()
	user := dbCfg.Key("USER").String()
	password := dbCfg.Key("PASSWORD").String()
	host := dbCfg.Key("HOST").String()
	charset := dbCfg.Key("CHARSET").String()
	parseTime := dbCfg.Key("PARSE_TIME").String()
	loc := dbCfg.Key("LOC").String()
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s", user, password, host, dbName, charset, parseTime, loc)
}


func init() {
	var err error
	
	dsn := getDSN()
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{ 
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

// 执行sql原语
func Exec(result interface{}, sql string, args ...interface{}) *gorm.DB {
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

func Count(model interface{}, count *int64) *gorm.DB {
	return db.Model(model).Count(count)
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
func Updates(updates interface{}) *gorm.DB {
	return db.Model(updates).Updates(updates)
}

// 删除
// 1. 用法：db.Where(条件表达式).Delete(空模型变量指针)
func Delete(model interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return db.Where(query, args).Delete(model)
}

func Scopes(fn func(db *gorm.DB) *gorm.DB) *gorm.DB {
	return db.Scopes(fn)
}

//分页封装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
