package service

import (
	"github.com/google/wire"
	v1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/usercenter/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserCenterService)

type UserCenterService struct {
	v1.UnimplementedUserCenterServer

	uc *biz.UserUseCase
}

func NewUserCenterService(uc *biz.UserUseCase) *UserCenterService {
	return &UserCenterService{uc: uc}
}
