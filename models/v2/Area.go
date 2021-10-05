package models_v2

type Area struct {
	Base
	Pid   string `json:"pid"`
	Title string `json:"title"`
	Code  string `json:"code"`
}
