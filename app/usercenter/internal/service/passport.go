package service

import (
	"context"
	v1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/pkg/errx"
)

// Register 注册用户
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

// PasswdLogin 用户密码登录
func (s *UserCenterService) PasswdLogin(ctx context.Context, req *v1.PasswdLoginReq) (resp *v1.PasswdLoginResp, err error) {
	u := &biz.User{
		Username: req.Username,
		AreaCode: req.AreaCode,
		Password: req.Password,
	}
	user, err := s.uc.GetMultiWayUser(ctx, u)
	if err != nil {
		return
	}
	if err = s.uc.VerifyUserPassport(ctx, user, u.Password); err != nil {
		return
	}
	_, err = s.uc.CreateUserToken(ctx, user)
	return
}
