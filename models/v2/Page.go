package models_v2

import (
	"goserver/libs/gorm"
	"time"
)

// 页面日志
type PageLog struct {
	Info      string    `json:"info"`
	Timestamp time.Time `json:"timestamp"`
}

// 页面错误
type PageError struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// 页面操作
type PageOperation struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// 页面信息
type Page struct {
	Base
	Title       string    `json:"title"`
	Icon        string    `json:"icon"`
	Keywords    FlatArray `gorm:"TYPE:json" json:"keywords"`
	Description string    `json:"description"`
	Path        string    `gorm:"not null" json:"path"`
	Query       FlatMap   `gorm:"TYPE:json" json:"query"`
	Template    int       `json:"template"`
	Width       int       `gorm:"default:1440" json:"width"`
	Height      int       `gorm:"default:1080" json:"height"`
	Tags        FlatArray `gorm:"TYPE:json" json:"tags"`
	Lang        string    `json:"lang"`
	Timezone    string    `json:"timezone"`
	Content     FlatArray `gorm:"TYPE:json" json:"content"`
	Logs        FlatArray `gorm:"TYPE:json" json:"logs"`
	Errors      FlatArray `gorm:"TYPE:json" json:"errors"`
	Operations  FlatArray `gorm:"TYPE:json" json:"operations"`
}

func init() {
	gorm.AutoMigrat(&Page{})
}
