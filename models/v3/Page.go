package models_v3

import (
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
	Title       string    `gorm:"TYPE:varchar(255)" json:"title"`
	Logo        string    `gorm:"TYPE:varchar(255)" json:"logo"`
	Background  FlatArray `gorm:"TYPE:json" json:"background"`
	Keywords    string    `gorm:"TYPE:varchar(255)" json:"keywords"`
	Description string    `gorm:"TYPE:varchar(255)" json:"description"`
	Path        string    `gorm:"TYPE:varchar(255);not null" json:"path"`
	Query       FlatArray `gorm:"TYPE:json" json:"query"`
	Template    string    `gorm:"TYPE:varchar(255);default:blank" json:"template"`
	Width       int       `gorm:"default:1440" json:"width"`
	Height      int       `gorm:"default:1080" json:"height"`
	Fullscreen  bool      `gorm:"default:false" json:"fullscreen"`
	Theme       string    `gorm:"TYPE:varchar(255);default:dark" json:"theme"`
	Tags        FlatArray `gorm:"TYPE:json" json:"tags"`
	Lang        string    `gorm:"TYPE:varchar(255)" json:"lang"`
	Timezone    string    `gorm:"TYPE:varchar(255)" json:"timezone"`
	Content     FlatMap   `gorm:"TYPE:json" json:"content"`
	Global      FlatArray `gorm:"TYPE:json" json:"global"`
	Logs        FlatArray `gorm:"TYPE:json" json:"logs"`
	Errors      FlatArray `gorm:"TYPE:json" json:"errors"`
	Operations  FlatArray `gorm:"TYPE:json" json:"operations"`
	ImgUrl      string    `gorm:"type:varchar(255);not null;column:img_url" json:"imgUrl"`
}
