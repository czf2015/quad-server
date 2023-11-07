package models_v3

import (
	"database/sql/driver"
	"encoding/json"
	"goserver/libs/utils"
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID         string     `gorm:"primaryKey" json:"id"`
	CreateTime *time.Time `gorm:"type:timestamp;default:NOW();autoCreateTime" json:"create_time"`
	UpdateTime *time.Time `gorm:"type:timestamp;default:NOW();autoUpdateTime" json:"update_time"`
	DeleteTime *time.Time `gorm:"type:timestamp;default:NULL" json:"-"`
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
