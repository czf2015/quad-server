package models_v2

type Area struct {
	Base
	Pid   string `json:"pid"`
	Title string `json:"label"`
	Code  string `json:"code"`
}
