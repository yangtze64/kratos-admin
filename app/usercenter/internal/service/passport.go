package service

import (
	"context"
	v1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/utils/errx"
)

func (s *UserCenterService) Register(ctx context.Context, req *v1.RegisterReq) (resp *v1.RegisterResp, err error) {
	if req.PasswordReview != req.Password {
		return nil, errx.New(errx.TowPasswordDiff)
	}
	u := &biz.User{
		Username: req.Username,
		Realname: req.Realname,
		Mobile:   req.Mobile,
		AreaCode: req.AreaCode,
		Password: req.Password,
		Email:    req.Email,
		Weixin:   req.Weixin,
		Operator: req.Operator,
	}
	user, err := s.uc.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	resp = &v1.RegisterResp{
		Username: user.Username,
		Mobile:   user.Mobile,
		AreaCode: user.AreaCode,
		Email:    user.Email,
	}
	return
}
