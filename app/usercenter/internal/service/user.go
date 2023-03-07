package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/utils/global"
	"time"
)

// CreateUser 创建用户
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
	return
}

func (s *UserCenterService) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (resp *emptypb.Empty, err error) {
	user := biz.User{
		Uid:       req.Uid,
		Username:  req.Username,
		Realname:  req.Realname,
		Mobile:    req.Mobile,
		AreaCode:  req.AreaCode,
		Email:     req.Email,
		Weixin:    req.Weixin,
		Operator:  req.Operator,
		UpdatedAt: req.UpdatedAt,
	}
	err = s.uc.UpdateUserByUid(ctx, &user)
	return
}

// DeleteUser 删除用户
func (s *UserCenterService) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (resp *emptypb.Empty, err error) {
	user := biz.User{
		Uid:       req.Uid,
		Operator:  req.Operator,
		DeletedAt: req.DeletedAt,
	}
	err = s.uc.DeleteUserByUid(ctx, &user)
	return
}

// FindUserByUid 根据UID获取用户
func (s *UserCenterService) FindUserByUid(ctx context.Context, req *v1.FindUserByUidReq) (resp *v1.User, err error) {
	u, err := s.uc.GetUserByUid(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	resp = UserFromBizUser(u)
	return
}

// ExistUserByUid 用户是否存在
func (s *UserCenterService) ExistUserByUid(ctx context.Context, req *v1.ExistUserByUidReq) (resp *v1.ExistUserByUidResp, err error) {
	var isExist bool
	isExist, err = s.uc.ExistUserByUid(ctx, req.Uid)
	resp = &v1.ExistUserByUidResp{
		IsExist: isExist,
	}
	return
}

func (s *UserCenterService) ListUser(ctx context.Context, req *v1.UserFilter) (resp *v1.ListUserResp, err error) {
	user := UserFilterToBizUser(req)
	data, total, err := s.uc.GetUserList(ctx, user)
	if err != nil {
		return nil, err
	}
	var list []*v1.User
	if n := len(data); n > 0 {
		list = make([]*v1.User, 0, n)
		for _, v := range data {
			list = append(list, UserFromBizUser(v))
		}
	}
	resp = &v1.ListUserResp{
		Total: total,
		List:  list,
	}
	return
}

func UserFromBizUser(u *biz.User) *v1.User {
	return &v1.User{
		Uid:         u.Uid,
		Username:    u.Username,
		Realname:    u.Realname,
		Mobile:      u.Mobile,
		AreaCode:    u.AreaCode,
		Email:       u.Email,
		Weixin:      u.Weixin,
		Operator:    u.Operator,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		CreatedDate: time.Unix(u.CreatedAt, 0).Format(global.TimeFormat),
		UpdatedDate: time.Unix(u.UpdatedAt, 0).Format(global.TimeFormat),
	}
}

func UserFilterToBizUser(req *v1.UserFilter) *biz.User {
	user := biz.User{
		Mobile:    req.Mobile,
		AreaCode:  req.AreaCode,
		Username:  req.Username,
		Realname:  req.Realname,
		Email:     req.Email,
		Weixin:    req.Weixin,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,

		Ids:              req.Id,
		Uids:             req.Uid,
		UsernameFuzzy:    req.UsernameFuzzy,
		RealnameFuzzy:    req.RealnameFuzzy,
		EmailFuzzy:       req.EmailFuzzy,
		Operators:        req.Operator,
		CreatedDateStart: req.CreatedDateStart,
		CreatedDateEnd:   req.CreatedDateEnd,
		UpdatedDateStart: req.UpdatedDateStart,
		UpdatedDateEnd:   req.UpdatedDateEnd,
		Page:             int(req.Page),
		Limit:            int(req.Limit),

		SortId:        int32(req.SortId),
		SortCreatedAt: int32(req.SortCreatedAt),
		SortUpdatedAt: int32(req.SortUpdatedAt),
	}
	if user.SortId == 0 && user.SortCreatedAt == 0 && user.SortUpdatedAt == 0 {
		user.SortId = int32(v1.UserFilter_DESC)
	}
	return &user
}
