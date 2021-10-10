package models_v2

type NetworkAllocation struct {
	Base
	UserId                  string `json:"user_id" form:"user_id"`
	Page int `json:"page" form:"page"`
	Bits                string `json:"bits"`
}
