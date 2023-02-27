package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Uid      string
	Username string
	Realname string
	mobile   string
	AreaCode string
	Email    string
	Weixin   string
	Unionid  string
}

type UserRepo interface {
	Create(ctx context.Context, user *User) (uid int64, err error)
	Update(ctx context.Context, uid string) (bool, error)
	Delete(ctx context.Context, uid string) (bool, error)
	List(ctx context.Context) ([]*User, error)
	FindByUid(ctx context.Context, uid string) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/user")),
	}
}
