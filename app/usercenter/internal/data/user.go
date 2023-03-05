package data

import (
	"context"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/app/usercenter/internal/model/sysuser"
	"kratos-admin/pkg/expr"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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
	m := sysuser.SysUser{
		Uid:      user.Uid,
		Username: user.Username,
		Password: user.Password,
		Realname: user.Realname,
		Mobile:   user.Mobile,
		AreaCode: user.AreaCode,
		Email:    user.Email,
		Weixin:   user.Weixin,
		Unionid:  user.Unionid,
		CreateAt: time.Unix(user.CreateTime,0).Local(),
		UpdateAt: time.Unix(user.UpdateTime,0).Local(),
		Operator: user.Operator,
	}

	err = u.data.db.WithContext(ctx).Create(&m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
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
	var user sysuser.SysUser
	err := u.data.db.WithContext(ctx).Omit(sysuser.Column.IsDelete.String(), sysuser.Column.DeleteAt.String()).
		Where(sysuser.Column.Uid.Eq(), uid).Limit(1).First(&user).Error
	if err != nil {
		return nil, err
	}
	return UserFromEntity(&user), nil
}
func (u *userRepo) ExistUsername(ctx context.Context, username string) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName).Select(sysuser.Column.Id.String()).
		Where(sysuser.Column.Username.Eq(), username)
	exist, err := u.exist(ctx, sub)
	return exist, err
}

func (u *userRepo) ExistMobile(ctx context.Context, mobile string, areaCode int32) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName).Select(sysuser.Column.Id.String()).
		Where(sysuser.Column.Mobile.Eq(),mobile).
		Where(sysuser.Column.AreaCode.Eq(),areaCode)
	exist, err := u.exist(ctx, sub)
	return exist, err
}
func (u *userRepo) ExistEmail(ctx context.Context, email string) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName).Select(sysuser.Column.Id.String()).
		Where(sysuser.Column.Email.Eq(),email)
	exist, err := u.exist(ctx, sub)
	return exist, err
}
func (u *userRepo) ExistUnionId(ctx context.Context, unionid string) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName).Select(sysuser.Column.Id.String()).
		Where(sysuser.Column.Unionid.Eq(),unionid)
	exist, err := u.exist(ctx, sub)
	return exist, err
}

func (u *userRepo) exist(ctx context.Context, sub *gorm.DB) (bool, error) {
	var exist bool
	sub.Where(sysuser.Column.IsDelete.Eq(), 0).Limit(1)
	// "IF(u.id > 0,1,0) as exist"
	st := sysuser.Column.Id.Expr(func(f expr.String) expr.String {
		return "IF(u.`" + f + "` > 0, " + expr.Symbol + ", " + expr.Symbol + ") as exist"
	}).String()
	err := u.data.db.WithContext(ctx).Select(st, 1, 0).Table("(?) as u", sub).Scan(&exist).Error
	if err != nil {
		return false, err
	}
	return exist, nil
}

func UserFromEntity(m *sysuser.SysUser) *biz.User {
	return &biz.User{
		Uid:      m.Uid,
		Username: m.Username,
		Realname: m.Realname,
		Mobile:   m.Mobile,
		AreaCode: m.AreaCode,
		Email:    m.Email,
		Weixin:   m.Weixin,
		Unionid:  m.Unionid,
		CreateTime: m.CreateAt.Unix(),
		UpdateTime: m.UpdateAt.Unix(),
		Operator: m.Operator,
	}
}
