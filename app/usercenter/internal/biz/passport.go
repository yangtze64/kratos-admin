package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type JwtToken struct {
	AccessToken  string
	AccessExpire int64
	RefreshAfter int64
}

type PassportRepo interface {
	Login(ctx context.Context, user *User) (*JwtToken, error)
	Logout(ctx context.Context) error
}

type PassportUseCase struct {
	repo PassportRepo
	log  *log.Helper
}

func NewPassportUseCase(repo PassportRepo, logger log.Logger) *PassportUseCase {
	return &PassportUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/passport")),
	}
}
