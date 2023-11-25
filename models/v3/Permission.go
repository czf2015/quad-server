package models_v3

type Permission struct {
	Base
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Path        string `gorm:"type:varchar(255);not null" json:"path"` // 页面路径
	Privilege   int    `json:"privilege"`                              // 位 从低到高依次表示查看、编辑、审核
	Description string `gorm:"type:varchar(255)" json:"description"`
}
