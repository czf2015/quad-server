package models_v2

import (
	"time"

	// "goserver/libs/gorm"
	// "goserver/libs/utils"
)

type Base struct {
	ID        string    `gorm:"primaryKey;" json:"id" form:"id"`
	CreatedAt time.Time `gorm:"default:NOW()" json:"-"`
	UpdatedAt time.Time `gorm:"default:NOW()" json:"-"`
	DeletedAt time.Time `gorm:"default:NULL" json:"-"`
}

// func (base *Base) BeforeCreate(db *gorm.DB) (err error) {
// 	if len(base.ID) == 0 {
//     base.ID = utils.GenerateUuid()
//   }

// 	return
// }
