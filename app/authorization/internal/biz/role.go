package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-admin/pkg/errx"
)

type Role struct {
	Id          int32
	Name        string
	Description string
	IsEnable    int32
	Operator    string
	CreatedAt   int32
	UpdatedAt   int32
}

type RoleRepo interface {
	Create(ctx context.Context, role *Role) (id int32, err error)
	ExistRoleName(ctx context.Context, name string, excludeIds ...int32) (exist bool, err error)
}

type RoleUseCase struct {
	repo RoleRepo
	log  *log.Helper
}

func NewRoleUseCase(repo RoleRepo, logger log.Logger) *RoleUseCase {
	return &RoleUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "authorization/usercase")),
	}
}

func (c *RoleUseCase) CreateRole(ctx context.Context, role *Role) (id int32, err error) {
	var exist bool
	if exist, err = c.repo.ExistRoleName(ctx, role.Name); err != nil {
		return
	}
	if exist {
		return 0, errx.New(errx.RoleNameExist)
	}
	id, err = c.repo.Create(ctx, role)
	return
}

func (c *RoleUseCase) UpdateRole(ctx context.Context, role *Role) (id int32, err error) {
	//var exist bool
	//if exist, err = c.repo.ExistRoleName(ctx, role.Name); err != nil {
	//	return
	//}
	//if exist {
	//	return 0, errx.New(errx.RoleNameExist)
	//}
	//id, err = c.repo.Create(ctx, role)
	return
}
