package models_v3

type Template struct {
	Page
	Remark string `gorm:"type:varchar(255)" json:"remark"`
}
