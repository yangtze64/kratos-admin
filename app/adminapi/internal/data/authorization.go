package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	authorizationv1 "kratos-admin/api/authorization/service/v1"
	"kratos-admin/app/adminapi/internal/conf"
)

type AuthorizationRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthorizationClient(sr *conf.Service, r registry.Discovery) authorizationv1.AuthorizationClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Authorization.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithTimeout(sr.Authorization.Timeout.AsDuration()),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := authorizationv1.NewAuthorizationClient(conn)
	return c
}

func NewAuthorizationRepo(data *Data, logger log.Logger) *AuthorizationRepo {
	return &AuthorizationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
