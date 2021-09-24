package models

type ApprovedDomain struct {
	Base

	UserId string `json:"id"`
	Domain string `json:"domain"`
	Approved bool `json:"approved"`
}