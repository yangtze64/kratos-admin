package utils

import "kratos-admin/pkg/global"

func WrapMobileAreaCode(areaCode int32) int32 {
	if areaCode == 0 {
		areaCode = global.DefaultMobileAreaCode
	}
	return areaCode
}
