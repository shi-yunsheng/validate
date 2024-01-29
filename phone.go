package validate

import "github.com/ttacon/libphonenumber"

// 是否是正确的电话号码
func IsPhone(phone string) bool {
	number, err := libphonenumber.Parse(phone, "")
	if err != nil {
		return false
	}
	if !libphonenumber.IsValidNumber(number) {
		return false
	}

	return true
}
