package data

import (
	"context"
	"kratos-admin/pkg/errx"
	"kratos-admin/pkg/global"
	"kratos-admin/pkg/hash"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/app/usercenter/internal/model/sysuser"
	"kratos-admin/pkg/expr"
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

func (u *userRepo) Create(ctx context.Context, user *biz.User) (id int32, err error) {
	m := sysuser.SysUser{
		Uid:       user.Uid,
		Username:  user.Username,
		Password:  user.Password,
		Realname:  user.Realname,
		Mobile:    user.Mobile,
		AreaCode:  user.AreaCode,
		Email:     user.Email,
		Weixin:    user.Weixin,
		Operator:  user.Operator,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	nowTime := time.Now().Unix()
	if m.CreatedAt <= 0 {
		m.CreatedAt = int32(nowTime)
	}
	if m.UpdatedAt <= 0 {
		m.UpdatedAt = int32(nowTime)
	}
	err = u.data.db.WithContext(ctx).Create(&m).Error
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}
func (u *userRepo) Update(ctx context.Context, uid string, user *biz.User) error {
	m := sysuser.SysUser{
		Username:  user.Username,
		Realname:  user.Realname,
		Mobile:    user.Mobile,
		AreaCode:  user.AreaCode,
		Email:     user.Email,
		Weixin:    user.Weixin,
		Operator:  user.Operator,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	if m.DeletedAt > 0 && m.UpdatedAt <= 0 {
		m.UpdatedAt = int32(time.Now().Unix())
	}
	err := u.data.db.WithContext(ctx).
		Where(sysuser.Column.Uid.Eq(), uid).
		Where(sysuser.Column.DeletedAt.Eq(), global.ModelNotDeleteAt).
		Updates(m).
		Error
	return err
}

func (u *userRepo) Delete(ctx context.Context, uid string, user *biz.User) error {
	resetUser := &biz.User{
		Operator:  user.Operator,
		DeletedAt: user.DeletedAt,
	}
	if resetUser.DeletedAt <= 0 {
		resetUser.DeletedAt = int32(time.Now().Unix())
	}
	return u.Update(ctx, uid, resetUser)
}

func (u *userRepo) List(ctx context.Context, user *biz.User) (list []*biz.User, total int64, err error) {
	userCondOptions := []UserCondOption{
		WithIdWhere(),
		WithUidWhere(),
		WithMobileWhere(),
		WithUsernameWhere(),
		WithUsernameFuzzyWhere(),
		WithEmailWhere(),
		WithEmailFuzzyWhere(),
		WithRealnameWhere(),
		WithRealnameFuzzyWhere(),
		WithWeixinWhere(),
		WithOperatorWhere(),
		WithCreatedAtWhere(),
		WithUpdatedAtWhere(),
	}
	tx := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName)
	if tx, err = UserCondChain(tx, user, userCondOptions...); err != nil {
		return
	}
	if err = tx.Count(&total).Error; err != nil {
		return
	}
	userCondOptions = append(userCondOptions, WithPager(), WithSortId(), WithSortCreatedAt(), WithSortUpdatedAt())
	sub := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName).Select(sysuser.Column.Id.String())
	if sub, err = UserCondChain(sub, user, userCondOptions...); err != nil {
		return
	}
	var users []*sysuser.SysUser
	query := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName+" AS t").Omit(sysuser.Column.DeletedAt.String(), sysuser.Column.Password.String()).
		InnerJoins("INNER JOIN (?) AS s ON t.id = s.id", sub)
	if query, err = UserCondChain(query, user, WithSortId(), WithSortCreatedAt(), WithSortUpdatedAt()); err != nil {
		return
	}
	if err = query.Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return list, total, nil
		}
		return
	}
	if len(users) > 0 {
		for _, v := range users {
			list = append(list, UserFromEntity(v))
		}
	}
	return
}
func (u *userRepo) FindByUid(ctx context.Context, uid string) (*biz.User, error) {
	var user sysuser.SysUser
	err := u.data.db.WithContext(ctx).Omit(sysuser.Column.DeletedAt.String()).
		Where(sysuser.Column.Uid.Eq(), uid).
		Where(sysuser.Column.DeletedAt.Eq(), global.ModelNotDeleteAt).
		Limit(1).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.New(errx.UserNotFound)
		}
		return nil, err
	}
	return UserFromEntity(&user), nil
}
func (u *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	var user sysuser.SysUser
	err := u.data.db.WithContext(ctx).Omit(sysuser.Column.DeletedAt.String()).
		Where(sysuser.Column.Username.Eq(), username).
		Where(sysuser.Column.DeletedAt.Eq(), global.ModelNotDeleteAt).
		Limit(1).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.New(errx.UserNotFound)
		}
		return nil, err
	}
	return UserFromEntity(&user), nil
}
func (u *userRepo) FindByMobile(ctx context.Context, mobile string, areaCode int32) (*biz.User, error) {
	var user sysuser.SysUser
	err := u.data.db.WithContext(ctx).Omit(sysuser.Column.DeletedAt.String()).
		Where(sysuser.Column.Mobile.Eq(), mobile).
		Where(sysuser.Column.AreaCode.Eq(), areaCode).
		Where(sysuser.Column.DeletedAt.Eq(), global.ModelNotDeleteAt).
		Limit(1).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.New(errx.UserNotFound)
		}
		return nil, err
	}
	return UserFromEntity(&user), nil
}
func (u *userRepo) FindByEmail(ctx context.Context, email string) (*biz.User, error) {
	var user sysuser.SysUser
	err := u.data.db.WithContext(ctx).Omit(sysuser.Column.DeletedAt.String()).
		Where(sysuser.Column.Email.Eq(), email).
		Where(sysuser.Column.DeletedAt.Eq(), global.ModelNotDeleteAt).
		Limit(1).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errx.New(errx.UserNotFound)
		}
		return nil, err
	}
	return UserFromEntity(&user), nil
}

func (u *userRepo) ExistUser(ctx context.Context, uid string) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName).Select(sysuser.Column.Id.String()).
		Where(sysuser.Column.Uid.Eq(), uid)
	exist, err := u.exist(ctx, sub)
	return exist, err
}
func (u *userRepo) ExistUsername(ctx context.Context, username string, excludeUids ...string) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName).Select(sysuser.Column.Id.String()).
		Where(sysuser.Column.Username.Eq(), username)
	if exn := len(excludeUids); exn > 0 {
		if exn == 1 {
			sub.Where(sysuser.Column.Uid.Neq(), excludeUids[0])
		} else {
			sub.Where(sysuser.Column.Uid.NotIn(), excludeUids)
		}
	}
	exist, err := u.exist(ctx, sub)
	return exist, err
}
func (u *userRepo) ExistMobile(ctx context.Context, mobile string, areaCode int32, excludeUids ...string) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName).Select(sysuser.Column.Id.String()).
		Where(sysuser.Column.Mobile.Eq(), mobile).
		Where(sysuser.Column.AreaCode.Eq(), areaCode)
	if exn := len(excludeUids); exn > 0 {
		if exn == 1 {
			sub.Where(sysuser.Column.Uid.Neq(), excludeUids[0])
		} else {
			sub.Where(sysuser.Column.Uid.NotIn(), excludeUids)
		}
	}
	exist, err := u.exist(ctx, sub)
	return exist, err
}
func (u *userRepo) ExistEmail(ctx context.Context, email string, excludeUids ...string) (bool, error) {
	sub := u.data.db.WithContext(ctx).Table(sysuser.TableSysUserName).Select(sysuser.Column.Id.String()).
		Where(sysuser.Column.Email.Eq(), email)
	if exn := len(excludeUids); exn > 0 {
		if exn == 1 {
			sub.Where(sysuser.Column.Uid.Neq(), excludeUids[0])
		} else {
			sub.Where(sysuser.Column.Uid.NotIn(), excludeUids)
		}
	}
	exist, err := u.exist(ctx, sub)
	return exist, err
}
func (u *userRepo) exist(ctx context.Context, sub *gorm.DB) (bool, error) {
	var exist bool
	sub.Where(sysuser.Column.DeletedAt.Eq(), global.ModelNotDeleteAt).Limit(1)
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

func (u *userRepo) CacheAccessToken(ctx context.Context, token string, expire int64) error {
	err := u.data.rds.Set(ctx, global.CacheUserLoginToken+hash.Md5Hex([]byte(token)), token, time.Second*time.Duration(expire)).Err()
	return err
}
func (u *userRepo) DelCacheAccessToken(ctx context.Context, token string) error {
	err := u.data.rds.Del(ctx, global.CacheUserLoginToken+hash.Md5Hex([]byte(token))).Err()
	return err
}

func UserFromEntity(m *sysuser.SysUser) *biz.User {
	return &biz.User{
		Id:        m.Id,
		Uid:       m.Uid,
		Username:  m.Username,
		Password:  m.Password,
		Realname:  m.Realname,
		Mobile:    m.Mobile,
		AreaCode:  m.AreaCode,
		Email:     m.Email,
		Weixin:    m.Weixin,
		Operator:  m.Operator,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
