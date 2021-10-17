package models_v2

type AddressPlan struct {
	Base
	UserId                  string `json:"user_id" form:"user_id"`
	NetworkAddress          string `json:"network_address" form:"network_address"`
	BitWidth                int    `json:"bit_width"`
	SubnetAddressBeginValue int    `json:"subnet_address_begin_value"`
	PrefixBitWidth          int    `json:"prefix_bit_width"`
	Organization            string `json:"organization" form:"organization"`
	AddressCount            int    `json:"address_count"`
	SubnetType              string `json:"subnet_type" form:"subnet_type"`
	AddressList             string `json:"address_list"`
}
