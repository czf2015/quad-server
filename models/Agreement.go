package models

import (
	"goserver/libs/db"
)

type Agreement struct {
    Base

    Type string `json:"type"`
    Content string `json:"content"`
}

func GetLatestAgreements() (latest []Agreement) {
    db.DB().Table("agreement").Where("created_at IN (SELECT MAX(created_at) FROM agreement WHERE deleted_at IS NULL GROUP BY type)").Find(&latest)
    return
}