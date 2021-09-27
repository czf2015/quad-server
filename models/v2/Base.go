package models_v2

import (
	"time"	     
)

type Base struct {
	ID        string    `gorm:"primaryKey" json:"id"`
  CreatedAt time.Time `gorm:"default:NOW()" json:"created_at"`
  UpdatedAt time.Time `gorm:"default:NOW()" json:"updated_at"`
  DeletedAt time.Time `gorm:"default:NULL" json:"deleted_at"`
}