package models_v2

import (
	"database/sql/driver"
	"encoding/json"
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
	Published   int       `gorm:"default:0" json:"published"`
	Version     string    `json:"version"`
	Content     FlatArray `gorm:"TYPE:json" json:"content"`
	Logs        FlatArray `gorm:"TYPE:json" json:"logs"`
	Errors      FlatArray `gorm:"TYPE:json" json:"errors"`
	Operations  FlatArray `gorm:"TYPE:json" json:"operations"`
}

func init() {
	gorm.AutoMigrat(&Page{})
}

type FlatMap map[string]interface{}

func (c FlatMap) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *FlatMap) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type FlatArray []interface{}

func (c FlatArray) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

// TODO: 标签查询存在问题
func (c *FlatArray) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// 作者：🐟本尊87045
// 链接：https://juejin.cn/post/6844904120516608008
// 来源：稀土掘金
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
