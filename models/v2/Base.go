package models_v2

import (
	"database/sql/driver"
	"encoding/json"
	"goserver/libs/gorm"
	"goserver/libs/utils"
	"time"
	// "goserver/libs/gorm"
	// "goserver/libs/utils"
)

type Base struct {
	ID        string     `gorm:"primaryKey;" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:NOW()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:NOW()" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp;default:NULL" json:"-"`
}

func (base *Base) BeforeCreate(db *gorm.DB) (err error) {
	if len(base.ID) == 0 {
		base.ID = utils.GenerateUuid()
	}

	return
}

type FlatMap map[string]interface{}

func (c FlatMap) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *FlatMap) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type FlatArray []interface{}

func (c FlatArray) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *FlatArray) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
