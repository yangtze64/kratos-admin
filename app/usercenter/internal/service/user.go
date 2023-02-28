package service

import (
	"context"
	v1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/usercenter/internal/biz"
)

func (s *UserCenterService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (resp *v1.CreateUserResp, err error) {
	u := &biz.User{
		Username: req.Username,
		Realname: req.Realname,
		Mobile:   req.Mobile,
		AreaCode: req.AreaCode,
		Password: req.Password,
		Email:    req.Email,
		Weixin:   req.Weixin,
		Operator: req.Operator,
		CreateAt: req.CreateAt.AsTime(),
		UpdateAt: req.UpdateAt.AsTime(),
	}

	user, err := s.uc.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	resp = &v1.CreateUserResp{
		Id:  int64(user.Id),
		Uid: user.Uid,
	}
	return &v1.CreateUserResp{}, nil
}
