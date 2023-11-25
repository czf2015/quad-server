package models_v3

type Menu struct {
	Base
	PID   uint   `json:"pid"`
	Title string `gorm:"type:varchar(255)" json:"title"`
	Link  string `gorm:"type:varchar(255)" json:"link"`
	Order int    `json:"order"`
}
