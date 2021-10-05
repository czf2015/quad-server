package models_v2

type AddressPlan struct {
	Base
	UserId                  string
	NetworkAddress          string
	BitWidth                int
	SubnetAddressBeginValue int
	PrefixBitWidth          int
	Organization            string
	AddressCount            int
	SubnetType              string
	AddressList             string
}
