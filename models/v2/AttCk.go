package models_v2


type AttCk struct {
	ID string
	Pid string 
	Type string // 战术 技术 子技术
	Name string
	Description string
	Url string // 链接
	Alerts int // 告警信息/条数
}
