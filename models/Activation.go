package models


type Activation struct {
	Base
	UserId string `json:"user_id"`
	CompletedAt string `gorm:"default:NULL" json:"completed_at"`
}