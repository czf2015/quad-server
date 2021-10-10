package models_v2

type NetworkManage struct {
	Base
	UserId                  string `json:"user_id" form:"user_id"`
	Organization            string `json:"organization" form:"organization"`
	SubnetType            string `json:"subnet_type" form:"subnet_type"`
	Subnet string `json:"subnet" form:"subnet"`
	Usage float64 `json:"usage"`
	Distributed float64 `json:"distributed"`
	CreateMethod            string `json:"create_method"`
}
