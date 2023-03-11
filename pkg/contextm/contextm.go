package contextm

import (
	"context"
	"kratos-admin/pkg/global"
)

func GetLoginUid(ctx context.Context) string {
	value, ok := ctx.Value(global.LoginUidKey).(string)
	if ok {
		return value
	}
	return ""
}

func GetLoginUsername(ctx context.Context) string {
	value, ok := ctx.Value(global.LoginUsernameKey).(string)
	if ok {
		return value
	}
	return ""
}

func GetLoginCurrToken(ctx context.Context) string {
	value, ok := ctx.Value(global.LoginCurrTokenKey).(string)
	if ok {
		return value
	}
	return ""
}
