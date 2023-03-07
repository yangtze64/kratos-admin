package utils

import "regexp"

func VerifyMobileFormat(mobile string) bool {
	rule := "^[0-9]{6,20}$"
	reg := regexp.MustCompile(rule)
	return reg.MatchString(mobile)
}

func VerifyEmailFormat(email string) bool {
	rule := "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	reg := regexp.MustCompile(rule)
	return reg.MatchString(email)
}
