package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-admin/app/usercenter/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) Create(ctx context.Context, user *biz.User) (uid int64, err error) {
	return 0, nil
}
func (u *userRepo) Update(ctx context.Context, uid string) (bool, error) {
	return false, nil
}
func (u *userRepo) Delete(ctx context.Context, uid string) (bool, error) {
	return false, nil
}
func (u *userRepo) List(ctx context.Context) ([]*biz.User, error) {
	return nil, nil
}
func (u *userRepo) FindByUid(ctx context.Context, uid string) (*biz.User, error) {
	return nil, nil
}
