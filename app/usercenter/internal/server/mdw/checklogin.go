package mdw

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	v1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/usercenter/internal/conf"
	"kratos-admin/app/usercenter/internal/data"
	"kratos-admin/pkg/errx"
	"kratos-admin/pkg/global"
	"kratos-admin/pkg/hash"
	"strings"
)

func WhiteListMatcher() selector.MatchFunc {
	WhiteList := map[string]struct{}{
		v1.OperationUserCenterRegister:    {},
		v1.OperationUserCenterPasswdLogin: {},
	}
	return func(ctx context.Context, operation string) bool {
		if _, ok := WhiteList[operation]; ok {
			return false
		}
		return true
	}
}

func setUserToCtx() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			claim, _ := jwt.FromContext(ctx)
			if claim == nil {
				return nil, errx.New(errx.UnauthorizedInfoMissing)
			}
			claimInfo := claim.(jwtv4.MapClaims)
			uid := claimInfo[global.LoginUidKey]
			usermame := claimInfo[global.LoginUsernameKey]
			if uid == "" {
				return nil, errx.New(errx.UnauthorizedInfoMissing)
			}
			ctx = context.WithValue(ctx, global.LoginUidKey, uid)
			ctx = context.WithValue(ctx, global.LoginUsernameKey, usermame)
			return handler(ctx, req)
		}
	}
}

func checkCacheToken() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				auths := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)
				token := auths[1]
				if tk := data.RdsCli.Get(ctx, global.CacheUserLoginToken+hash.Md5Hex([]byte(token))).Val(); tk == "" {
					return nil, errx.New(errx.UnauthorizedTokenInvalid)
				}
				ctx = context.WithValue(ctx, global.LoginCurrTokenKey, token)
			}
			return handler(ctx, req)
		}
	}
}

func CheckLogin(jwtconf *conf.JwtAuth) middleware.Middleware {
	return selector.Server(jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
		return []byte(jwtconf.Secret), nil
	}), checkCacheToken(), setUserToCtx()).Match(WhiteListMatcher()).Build()
}
