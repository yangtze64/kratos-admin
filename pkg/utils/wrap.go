package utils

import (
	"context"
	"kratos-admin/pkg/contextm"
	"kratos-admin/pkg/global"
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
