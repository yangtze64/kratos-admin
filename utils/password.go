package utils

import "kratos-admin/pkg/hash"

func GenPasswd(str string) string {
	return hash.Md5Hex([]byte(str))
}
