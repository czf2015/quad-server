package models_v3

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Base struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" form:"id" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp;default:NOW();autoCreateTime" json:"create_time"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:NOW();autoUpdateTime" json:"update_time"`
	DeletedAt *time.Time `gorm:"type:timestamp;default:NULL" json:"-"`
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
