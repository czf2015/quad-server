package models_v1

type ApprovedDomain struct {
	Base

	UserId string `json:"user_id"`
	Domain string `json:"domain"`
	Approved bool `json:"approved"`
}