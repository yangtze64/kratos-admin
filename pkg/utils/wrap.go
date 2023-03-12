package utils

import (
	"context"
	"kratos-admin/pkg/contextm"
	"kratos-admin/pkg/global"
	"strings"
)

func WrapMobileAreaCode(areaCode int32) int32 {
	if areaCode == 0 {
		areaCode = global.DefaultMobileAreaCode
	}
	return areaCode
}

func WrapOperator(ctx context.Context, operator string) string {
	if operator == "" {
		return contextm.GetLoginUid(ctx)
	}
	return operator
}

func WrapSensitiveStr(str string) (result string) {
	if str == "" {
		return "***"
	}
	if VerifyEmailFormat(str) {
		// email
		arr := strings.SplitN(str, "@", 2)
		s := []rune(arr[0])
		if len(s) <= 3 {
			result = "***@" + arr[1]
		} else {
			result = string(s[:3]) + "***@" + arr[1]
		}
	} else if VerifyMobileFormat(str) {
		// mobile
		strLen := len(str)
		if strLen <= 7 {
			result = str[:2] + "****" + str[strLen-2:strLen]
		} else {
			result = str[:3] + "****" + str[strLen-4:strLen]
		}
	} else {
		s := []rune(str)
		sLen := len(s)
		if sLen < 3 {
			result = string(s[:1]) + "**"
		} else if sLen >= 3 && sLen <= 5 {
			result = string(s[:1]) + "*" + string(s[sLen-1:sLen])
		} else {
			result = string(s[:2]) + "**" + string(s[sLen-2:sLen])
		}
	}
	return
}
