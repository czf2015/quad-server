package models_v2

import (
	"time"	     
)

type  Base struct {
	ID        string    `gorm:"primaryKey" json:"id"`
  CreatedAt time.Time `gorm:"default:NOW()" json:"-"`
  UpdatedAt time.Time `gorm:"default:NOW()" json:"-"`
  DeletedAt time.Time `gorm:"default:NULL" json:"-"`
}
