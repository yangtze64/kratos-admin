package service

import (
	v1 "kratos-admin/api/authorization/service/v1"
	"kratos-admin/app/authorization/internal/biz"
)

type AuthorizationService struct {
	v1.UnimplementedAuthorizationServer
	rc *biz.RoleUseCase
}

func NewAuthorizationService(rc *biz.RoleUseCase) *AuthorizationService {
	return &AuthorizationService{rc: rc}
}
