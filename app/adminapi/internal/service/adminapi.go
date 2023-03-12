package service

import (
	v1 "kratos-admin/api/adminapi/service/v1"
	"kratos-admin/app/adminapi/internal/data"
)

type AdminApiService struct {
	v1.UnimplementedAdminApiServer
	userCenterRepo    *data.UserCenterRepo
	authorizationRepo *data.AuthorizationRepo
}

func NewAdminApiService(userCenterRepo *data.UserCenterRepo, authorizationRepo *data.AuthorizationRepo) *AdminApiService {
	return &AdminApiService{userCenterRepo: userCenterRepo, authorizationRepo: authorizationRepo}
}
