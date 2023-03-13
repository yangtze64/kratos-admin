package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "kratos-admin/api/adminapi/service/v1"
	userCenterv1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/adminapi/internal/conf"
)

type UserCenterRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserCenterClient(sr *conf.Service, r registry.Discovery) userCenterv1.UserCenterClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(sr.Usercenter.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithTimeout(sr.Usercenter.Timeout.AsDuration()),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := userCenterv1.NewUserCenterClient(conn)
	return c
}

func NewUserCenterRepo(data *Data, logger log.Logger) *UserCenterRepo {
	return &UserCenterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *UserCenterRepo) Login(ctx context.Context, req *v1.LoginReq) (resp *v1.LoginResp, err error) {
	reply, err := repo.data.UserCenterClient.PasswdLogin(ctx, &userCenterv1.PasswdLoginReq{
		Username: req.Username,
		AreaCode: req.AreaCode,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	resp = &v1.LoginResp{
		Uid:          reply.Uid,
		Username:     reply.Username,
		AccessToken:  reply.AccessToken,
		AccessExpire: reply.AccessExpire,
		RefreshAfter: reply.RefreshAfter,
	}
	return
}
