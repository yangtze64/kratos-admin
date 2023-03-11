package service

import (
	"context"
	v1 "kratos-admin/api/authorization/service/v1"
	"kratos-admin/app/authorization/internal/biz"
	"kratos-admin/pkg/utils"
)

func (s *AuthorizationService) CreateRole(ctx context.Context, req *v1.CreateRoleReq) (resp *v1.CreateRoleResp, err error) {
	role := &biz.Role{
		Name:        req.Name,
		Description: req.Description,
		IsEnable:    req.IsEnable,
		Operator:    utils.WrapOperator(ctx, req.Operator),
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}
	id, err := s.rc.CreateRole(ctx, role)
	if err != nil {
		return nil, err
	}
	resp = &v1.CreateRoleResp{
		Id: id,
	}
	return
}
