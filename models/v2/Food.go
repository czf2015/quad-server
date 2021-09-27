package models_v2

import (
	"time"
)

//商品
type Food struct {
	Id         uint
	Title      string
	Price      float32
	Stock      int
	Type       int
	//mysql datetime, date类型字段，可以和golang time.Time类型绑定， 详细说明请参考：gorm连接数据库章节。
	CreateTime time.Time
}

//为Food绑定表名
func (v Food) TableName() string {
	return "foods"
}