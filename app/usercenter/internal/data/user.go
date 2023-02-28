package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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
func (u *userRepo) FindByUid(ctx context.Context, uid string) (*biz.User, error) {
	m := u.data.um.SysUser
	userEntity, err := m.WithContext(ctx).Omit(m.ID, m.IsDeleted, m.DeletedAt).
		Where(m.UID.Eq(uid)).Where(m.IsDeleted.Is(false)).First()
	if err != nil {
		return nil, err
	}
	return UserFromUserEntity(userEntity), nil
}
func (u *userRepo) ExistUsername(ctx context.Context, username string) (bool, error) {
	m := u.data.um.SysUser
	sub := m.WithContext(ctx).Select(m.ID).Where(m.Username.Eq(username)).Where(m.IsDeleted.Is(false))
	exist, err := u.exist(ctx, sub)
	return exist, err
}

func (u *userRepo) ExistMobile(ctx context.Context, mobile string, areaCode int32) (bool, error) {
	m := u.data.um.SysUser
	sub := m.WithContext(ctx).Select(m.ID).Where(m.Mobile.Eq(mobile)).Where(m.AreaCode.Eq(areaCode)).Where(m.IsDeleted.Is(false))
	exist, err := u.exist(ctx, sub)
	return exist, err
}
func (u *userRepo) ExistEmail(ctx context.Context, email string) (bool, error) {
	m := u.data.um.SysUser
	sub := m.WithContext(ctx).Select(m.ID).Where(m.Email.Eq(email)).Where(m.IsDeleted.Is(false))
	exist, err := u.exist(ctx, sub)
	return exist, err
}
func (u *userRepo) ExistUnionId(ctx context.Context, unionid string) (bool, error) {
	m := u.data.um.SysUser
	sub := m.WithContext(ctx).Select(m.ID).Where(m.Unionid.Eq(unionid)).Where(m.IsDeleted.Is(false))
	exist, err := u.exist(ctx, sub)
	return exist, err
}

func (u *userRepo) exist(ctx context.Context, sub interface{ UnderlyingDB() *gorm.DB }) (bool, error) {
	var exist bool
	m := u.data.um.SysUser
	err := m.WithContext(ctx).Exists(sub).Scan(&exist)
	if err != nil {
		return false, nil
	}
	return exist, nil
}

func UserFromUserEntity(m *entity.SysUser) *biz.User {
	return &biz.User{
		Uid:       m.UID,
		Username:  m.Username,
		Realname:  m.Realname,
		Mobile:    m.Mobile,
		AreaCode:  m.AreaCode,
		Email:     m.Email,
		Weixin:    m.Weixin,
		Unionid:   m.Unionid,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Operator:  m.Operator,
	}
}
