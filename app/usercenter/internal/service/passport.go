package service

import (
	"context"
	v1 "kratos-admin/api/usercenter/service/v1"
)

func (s *UserCenterService) Register(context.Context, *v1.RegisterReq) (*v1.RegisterResp, error) {

	return &v1.RegisterResp{}, nil
}
