package utils

import "kratos-admin/pkg/hash"

func GenPasswd(str string) string {
	return hash.Md5Hex([]byte(str))
}

func VerifyPassword(pass, str string) bool {
	if pass == GenPasswd(str) {
		return true
	}
	return false
}
