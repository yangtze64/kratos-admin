package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/app/usercenter/internal/model"
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

func (u *userRepo) Create(ctx context.Context, user *biz.User) (id int, err error) {
	entity := model.SysUser{
		UID:      user.Uid,
		Username: user.Username,
		Password: user.Password,
		Realname: user.Realname,
		Mobile:   user.Mobile,
		AreaCode: user.AreaCode,
		Email:    user.Email,
		Weixin:   user.Weixin,
		Unionid:  user.Unionid,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
		Operator: user.Operator,
	}
	err = u.data.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return 0, err
	}
	return entity.ID, nil
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
func (u *userRepo) ExistUsername(ctx context.Context, username string) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(model.TableSysUserName).Select("id").
		Where("username = ?", username)
	exist, err := u.exist(ctx, sub)
	return exist, err
}

func (u *userRepo) ExistMobile(ctx context.Context, mobile string, areaCode int32) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(model.TableSysUserName).Select("id").
		Where("mobile = ? AND area_code = ?", mobile, areaCode)
	exist, err := u.exist(ctx, sub)
	return exist, err
}
func (u *userRepo) ExistEmail(ctx context.Context, email string) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(model.TableSysUserName).Select("id").
		Where("email = ?", email)
	exist, err := u.exist(ctx, sub)
	return exist, err
}
func (u *userRepo) ExistUnionId(ctx context.Context, unionid string) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(model.TableSysUserName).Select("id").
		Where("unionid = ?", unionid)
	exist, err := u.exist(ctx, sub)
	return exist, err
}

func (u *userRepo) exist(ctx context.Context, sub *gorm.DB) (bool, error) {
	var exist bool
	sub.Where("is_delete = 0").Limit(1)
	err := u.data.db.WithContext(ctx).Select("IF(u.id > 0,1,0) as exist").Table("(?) as u", sub).Scan(&exist).Error
	if err != nil {
		return false, nil
	}
	return exist, nil
}

func UserFromUserEntity(m *model.SysUser) *biz.User {
	return &biz.User{
		Uid:      m.UID,
		Username: m.Username,
		Realname: m.Realname,
		Mobile:   m.Mobile,
		AreaCode: m.AreaCode,
		Email:    m.Email,
		Weixin:   m.Weixin,
		Unionid:  m.Unionid,
		CreateAt: m.CreateAt,
		UpdateAt: m.UpdateAt,
		Operator: m.Operator,
	}
}
