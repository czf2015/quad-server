package models_v2

type AddressPlan struct {
	Base
	UserId                  string `form:"userId"`
	NetworkAddress          string `form:"networkAddress"`
	BitWidth                int
	SubnetAddressBeginValue int
	PrefixBitWidth          int
	Organization            string `form:"organization"`
	AddressCount            int
	SubnetType              string `form:"subnetType"`
	AddressList             string
}
