package models_v2

import (
	"goserver/libs/gorm"
	"goserver/libs/utils"
	"time"
	// "goserver/libs/gorm"
	// "goserver/libs/utils"
)

type Base struct {
	ID        string    `gorm:"primaryKey;" json:"id"`
	CreatedAt time.Time `gorm:"type:timestamp;default:NOW()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:NOW()" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:timestamp;default:NULL" json:"-"`
}

func (base *Base) BeforeCreate(db *gorm.DB) (err error) {
	if len(base.ID) == 0 {
		base.ID = utils.GenerateUuid()
	}

	return
}
