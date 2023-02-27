package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-admin/app/usercenter/internal/biz"
)

type passportRepo struct {
	data *Data
	log  *log.Helper
}

func NewPassportRepo(data *Data, logger log.Logger) biz.PassportRepo {
	return &passportRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (p *passportRepo) Login(ctx context.Context, user *biz.User) (*biz.JwtToken, error) {
	return nil, nil
}
func (p *passportRepo) Logout(ctx context.Context) error {
	return nil
}
