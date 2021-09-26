package models2

type Base struct {
	ID         string `gorm:"primary_key" json:"id"`
	CreatedAt  string `gorm:"default:NOW()" json:"created_at"`
	ModifiedAt string `gorm:"default:NOW()" json:"modified_at"`
}