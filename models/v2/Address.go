package models_v2

type Address struct {
	Base
	Pid string `json:"pid"`
	Title string `gorm:"column:title" json:"title"`
	Link string `json:"link"`
	Icon string `json:"icon"`
}