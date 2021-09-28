package models_v1

type Base struct {
	ID         string `gorm:"primary_key" json:"id"`
	CreatedAt  string `gorm:"default:NOW()" json:"-"`
	ModifiedAt string `gorm:"default:NOW()" json:"-"`
}
