package service

import (
	"context"
	"gorm.io/gorm"
	v1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/utils/errx"
	"time"
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
		CreateTime:  req.CreateTime,
		UpdateTime:  req.UpdateTime,
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

func (s *UserCenterService) FindUserByUid(ctx context.Context, req *v1.FindUserByUidReq) (resp *v1.User, err error) {
	u, err := s.uc.GetUserByUid(ctx, req.Uid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.New(errx.UserNotFound)
		}
		return nil, err
	}
	return &v1.User{
		Uid:      u.Uid,
		Username: u.Username,
		Realname: u.Realname,
		Mobile:   u.Mobile,
		AreaCode: u.AreaCode,
		Email:    u.Email,
		Weixin:   u.Weixin,
		Unionid:  u.Unionid,
		CreateTime: u.CreateTime,
		UpdateTime: u.UpdateTime,
		CreateAt: time.Unix(u.CreateTime,0).Local().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Unix(u.UpdateTime,0).Local().Format("2006-01-02 15:04:05"),
		Operator: u.Operator,
	}, nil
}


