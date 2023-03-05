package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/utils/global"
	"time"
)

func (s *UserCenterService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (resp *v1.CreateUserResp, err error) {
	u := &biz.User{
		Username:  req.Username,
		Realname:  req.Realname,
		Mobile:    req.Mobile,
		AreaCode:  req.AreaCode,
		Password:  req.Password,
		Email:     req.Email,
		Weixin:    req.Weixin,
		Operator:  req.Operator,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
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

func (s *UserCenterService) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (resp *emptypb.Empty, err error) {
	user := biz.User{
		Uid: req.Uid,
	}
	if req.Operator != "" {
		user.Operator = req.Operator
	}
	if req.DeletedAt <= 0 {
		user.DeletedAt = time.Now().Unix()
	} else {
		user.DeletedAt = req.DeletedAt
	}
	err = s.uc.DeleteUserByUid(ctx, &user)
	return
}

func (s *UserCenterService) FindUserByUid(ctx context.Context, req *v1.FindUserByUidReq) (resp *v1.User, err error) {
	u, err := s.uc.GetUserByUid(ctx, req.Uid)
	if err != nil {
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
		Operator: u.Operator,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		CreatedDate: time.Unix(u.CreatedAt,0).Format(global.TimeFormat),
		UpdatedDate: time.Unix(u.UpdatedAt,0).Format(global.TimeFormat),

	}, nil
}



