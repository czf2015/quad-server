package models_v3

type Role struct {
	Base
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"`
	Permissions FlatArray `json:"permissions"`
}
