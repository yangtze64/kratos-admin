package service

import (
	"context"
	v1 "kratos-admin/api/usercenter/service/v1"
)

func (s *UserCenterService) CreateUser(context.Context, *v1.CreateUserReq) (*v1.CreateUserResp, error) {

	return &v1.CreateUserResp{}, nil
}
