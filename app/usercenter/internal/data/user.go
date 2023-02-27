package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/app/usercenter/internal/model/entity"
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
func (u *userRepo) FindByUid(ctx context.Context, uid string) (user *entity.SysUser, err error) {
	m := u.data.user.SysUser
	user, err = m.WithContext(ctx).
		Select(m.UID, m.Username, m.Realname, m.Mobile, m.AreaCode, m.Email, m.Weixin, m.Unionid, m.CreatedAt, m.UpdatedAt, m.Operator).
		Where(m.UID.Eq(uid)).Where(m.IsDeleted.Is(false)).First()
	if err != nil {
		return nil, err
	}
	return
}

//func UserFromSysUserEntity(entity *entity.SysUser, user *biz.User) {
//	user.Uid = entity.UID
//	user.Username = entity.Username
//	user.Realname = entity.Realname
//	user.Mobile = entity.Mobile
//	user.AreaCode = entity.AreaCode
//	user.Email = entity.Email
//	user.Weixin = entity.Weixin
//	user.Unionid = entity.Unionid
//}
