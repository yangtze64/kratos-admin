package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-admin/pkg/errx"
)

type Role struct {
	Id          int
	Name        string
	Description string
	IsEnable    bool
	Operator    string
	CreatedAt   int
	UpdatedAt   int
}

type RoleRepo interface {
	Create(ctx context.Context, role *Role) (id int, err error)
	ExistRoleName(ctx context.Context, name string) (exist bool, err error)
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

func (c *RoleUseCase) CreateRole(ctx context.Context, role *Role) (id int, err error) {
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
