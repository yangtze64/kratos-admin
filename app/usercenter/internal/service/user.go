package service

import (
	"context"
	v1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/usercenter/internal/biz"
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

//func (s *UserCenterService) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (resp *emptypb.Empty, err error) {
//	data := map[string]interface{}{
//		"uid":      req.Uid,
//	}
//	if req.Operator != "" {
//		data["operator"] = req.Operator
//	}
//	if req.DeleteTime <= 0 {
//		data["delete_at"] = time.Now()
//	} else {
//		data["delete_at"] = time.Unix(req.DeleteTime, 0)
//	}
//	err = s.uc.DeleteUserByUid(ctx, data)
//	return
//}

//func (s *UserCenterService) FindUserByUid(ctx context.Context, req *v1.FindUserByUidReq) (resp *v1.User, err error) {
//	u, err := s.uc.GetUserByUid(ctx, req.Uid)
//	if err != nil {
//		return nil, err
//	}
//	return &v1.User{
//		Uid:      u.Uid,
//		Username: u.Username,
//		Realname: u.Realname,
//		Mobile:   u.Mobile,
//		AreaCode: u.AreaCode,
//		Email:    u.Email,
//		Weixin:   u.Weixin,
//		Unionid:  u.Unionid,
//		CreateTime: u.CreateTime,
//		UpdateTime: u.UpdateTime,
//		CreateAt: time.Unix(u.CreateTime,0).Format(global.TimeFormat),
//		UpdateAt: time.Unix(u.UpdateTime,0).Format(global.TimeFormat),
//		Operator: u.Operator,
//	}, nil
//}



