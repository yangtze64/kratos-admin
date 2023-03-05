package service

import (
	"context"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	v1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/utils/errx"
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
		CreateAt: timestamppb.New(u.CreateAt),
		UpdateAt: timestamppb.New(u.UpdateAt),
		Operator: u.Operator,
	}, nil
}


