package service

import (
	"context"
	v1 "kratos-admin/api/usercenter/service/v1"
)

func (s *UserCenterService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (resp *v1.CreateUserResp, err error) {
	return &v1.CreateUserResp{}, nil
}
