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
func (s *UserCenterService) PasswdLogin(ctx context.Context, req *v1.PasswdLoginReq) (resp *v1.UserLoginResp, err error) {
	ud := &biz.User{
		Username: req.Username,
		AreaCode: req.AreaCode,
		Password: req.Password,
	}
	user, token, err := s.uc.UserPasswdLogin(ctx, ud)
	if err != nil {
		return
	}
	resp = &v1.UserLoginResp{
		Uid:          user.Uid,
		Username:     user.Username,
		Realname:     user.Realname,
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
		RefreshAfter: token.RefreshAfter,
	}
	return
}

// SimulationLogin 模拟登录
func (s *UserCenterService) SimulationLogin(ctx context.Context, req *v1.SimulationLoginReq) (resp *v1.UserLoginResp, err error) {
	ud := &biz.User{
		Mobile:   req.Mobile,
		AreaCode: req.AreaCode,
	}
	user, token, err := s.uc.UserMobileSimulationLogin(ctx, ud)
	if err != nil {
		return
	}
	resp = &v1.UserLoginResp{
		Uid:          user.Uid,
		Username:     user.Username,
		Realname:     user.Realname,
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
		RefreshAfter: token.RefreshAfter,
	}
	return
}
