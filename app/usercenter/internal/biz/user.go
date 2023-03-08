package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-admin/app/usercenter/internal/conf"
	"kratos-admin/pkg/errx"
	"kratos-admin/pkg/global"
	"kratos-admin/pkg/jwt"
	"kratos-admin/pkg/utils"
	"time"
)

type User struct {
	Id        int
	Uid       string
	Username  string
	Realname  string
	Password  string
	Mobile    string
	AreaCode  int32
	Email     string
	Weixin    string
	Operator  string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt int64

	Ids              []int64
	Uids             []string
	UsernameFuzzy    string
	RealnameFuzzy    string
	EmailFuzzy       string
	Operators        []string
	CreatedDateStart string
	CreatedDateEnd   string
	UpdatedDateStart string
	UpdatedDateEnd   string

	Page  int
	Limit int

	SortId        int32
	SortCreatedAt int32
	SortUpdatedAt int32
}

type JwtToken struct {
	AccessToken  string
	AccessExpire int64
	RefreshAfter int64
}

type UserRepo interface {
	Create(ctx context.Context, user *User) (id int, err error)
	Update(ctx context.Context, uid string, user *User) error
	List(ctx context.Context, user *User) (list []*User, total int64, err error)
	FindByUid(ctx context.Context, uid string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByMobile(ctx context.Context, mobile string, areaCode int32) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	ExistUser(ctx context.Context, uid string) (bool, error)
	ExistUsername(ctx context.Context, username string, excludeUids ...string) (bool, error)
	ExistMobile(ctx context.Context, mobile string, areaCode int32, excludeUids ...string) (bool, error)
	ExistEmail(ctx context.Context, email string, excludeUids ...string) (bool, error)

	CacheAccessToken(ctx context.Context, token string, expire int64) error
}

type UserUseCase struct {
	repo    UserRepo
	jwtAuth *conf.JwtAuth
	log     *log.Helper
}

func NewUserUseCase(repo UserRepo, jwtAuth *conf.JwtAuth, logger log.Logger) *UserUseCase {
	fmt.Printf("%+v\n", jwtAuth)
	return &UserUseCase{
		repo:    repo,
		jwtAuth: jwtAuth,
		log:     log.NewHelper(log.With(logger, "module", "usecase/user")),
	}
}

func (u *UserUseCase) UserMobileSimulationLogin(ctx context.Context, ud *User) (user *User, token *JwtToken, err error) {
	if ud.AreaCode == 0 {
		ud.AreaCode = global.DefaultMobileAreaCode
	}
	user, err = u.repo.FindByMobile(ctx, ud.Mobile, ud.AreaCode)
	if err != nil {
		return nil, nil, err
	}
	token, err = u.CreateUserToken(ctx, user)
	if err != nil {
		return nil, nil, err
	}
	return
}

func (u *UserUseCase) UserPasswdLogin(ctx context.Context, ud *User) (user *User, token *JwtToken, err error) {
	user, err = u.GetMultiWayUser(ctx, ud)
	if err != nil {
		return nil, nil, err
	}
	if err = u.VerifyUserPassport(ctx, user, ud.Password); err != nil {
		return nil, nil, err
	}
	token, err = u.CreateUserToken(ctx, user)
	if err != nil {
		return nil, nil, err
	}
	return
}

// CreateUserToken 创建用户Token
func (u *UserUseCase) CreateUserToken(ctx context.Context, user *User) (*JwtToken, error) {
	now := time.Now().Unix()
	uid := user.Uid
	expire := u.jwtAuth.Expire.Seconds
	payloads := map[string]interface{}{
		"iss":              u.jwtAuth.Issuer,
		"jti":              u.jwtAuth.Id,
		"sub":              uid,
		global.LoginUidKey: uid,
	}
	token, err := jwt.GenToken(now, u.jwtAuth.Secret, payloads, expire)
	if err != nil {
		return nil, err
	}
	// token缓存
	if err = u.repo.CacheAccessToken(ctx, token, expire); err != nil {
		return nil, err
	}
	return &JwtToken{
		AccessToken:  token,
		AccessExpire: expire,
		RefreshAfter: expire / 2,
	}, nil
}

// VerifyUserPassport 校验用户密码
func (u *UserUseCase) VerifyUserPassport(ctx context.Context, user *User, password string) error {
	if !utils.VerifyPassword(user.Password, password) {
		return errx.New(errx.UsernameOrPasswordIncorrect)
	}
	return nil
}

// GetMultiWayUser 获取多种方式得到用户
func (u *UserUseCase) GetMultiWayUser(ctx context.Context, user *User) (*User, error) {
	var (
		err     error
		info    *User
		isExist = false
	)
	if !isExist {
		if info, err = u.repo.FindByUsername(ctx, user.Username); err != nil {
			if !errors.Is(err, errx.New(errx.UserNotFound)) {
				return nil, err
			}
		}
		if info != nil {
			isExist = true
		}
	}
	if !isExist && utils.VerifyMobileFormat(user.Username) {
		if user.AreaCode == 0 {
			user.AreaCode = global.DefaultMobileAreaCode
		}
		if info, err = u.repo.FindByMobile(ctx, user.Username, user.AreaCode); err != nil {
			if !errors.Is(err, errx.New(errx.UserNotFound)) {
				return nil, err
			}
		}
		if info != nil {
			isExist = true
		}
	}
	if !isExist && utils.VerifyEmailFormat(user.Username) {
		if info, err = u.repo.FindByEmail(ctx, user.Username); err != nil {
			if !errors.Is(err, errx.New(errx.UserNotFound)) {
				return nil, err
			}
			isExist = false
		}
		if info != nil {
			isExist = true
		}
	}
	if !isExist {
		return nil, errx.New(errx.UserNotFound)
	}
	return info, nil
}

// CreateUser 创建用户
func (u *UserUseCase) CreateUser(ctx context.Context, user *User) (*User, error) {
	checkOptions := []UserCheckOption{
		WithExistUsername(u.repo),
		WithExistMobile(u.repo),
	}
	if user.Email != "" {
		checkOptions = append(checkOptions, WithExistEmail(u.repo))
	}
	if user.AreaCode == 0 {
		user.AreaCode = global.DefaultMobileAreaCode
	}
	err := UserCheckChain(ctx, user, checkOptions...)
	if err != nil {
		return nil, err
	}
	user.Uid = utils.NewUuid()
	user.Password = utils.GenPasswd(user.Password)
	nowTime := time.Now().Unix()
	if user.CreatedAt <= 0 {
		user.CreatedAt = nowTime
	}
	if user.UpdatedAt <= 0 {
		user.UpdatedAt = nowTime
	}
	id, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = id
	return user, nil
}

// UpdateUserByUid 根据UID更新用户
func (u *UserUseCase) UpdateUserByUid(ctx context.Context, user *User) error {
	checkOptions := []UserCheckOption{
		WithNotExistUser(u.repo),
	}
	if user.Username != "" {
		checkOptions = append(checkOptions, WithExistUsername(u.repo, user.Uid))
	}
	if user.AreaCode == 0 {
		user.AreaCode = global.DefaultMobileAreaCode
	}
	if user.Mobile != "" {
		checkOptions = append(checkOptions, WithExistMobile(u.repo, user.Uid))
	}
	if user.Email != "" {
		checkOptions = append(checkOptions, WithExistEmail(u.repo, user.Uid))
	}
	if err := UserCheckChain(ctx, user, checkOptions...); err != nil {
		return err
	}
	if user.UpdatedAt <= 0 {
		user.UpdatedAt = time.Now().Unix()
	}
	uid := user.Uid
	user.Uid = ""
	return u.repo.Update(ctx, uid, user)
}

// DeleteUserByUid 根据UID删除用户
func (u *UserUseCase) DeleteUserByUid(ctx context.Context, user *User) error {
	if err := UserCheckChain(ctx, user, WithNotExistUser(u.repo)); err != nil {
		return err
	}
	if user.DeletedAt <= 0 {
		user.DeletedAt = time.Now().Unix()
	}
	uid := user.Uid
	user.Uid = ""
	return u.repo.Update(ctx, uid, user)
}

// GetUserByUid 根据UID获取用户详情
func (u *UserUseCase) GetUserByUid(ctx context.Context, uid string) (user *User, err error) {
	user, err = u.repo.FindByUid(ctx, uid)
	return
}
func (u *UserUseCase) ExistUserByUid(ctx context.Context, uid string) (exist bool, err error) {
	exist, err = u.repo.ExistUser(ctx, uid)
	return
}
func (u *UserUseCase) GetUserList(ctx context.Context, user *User) (list []*User, total int64, err error) {
	list, total, err = u.repo.List(ctx, user)
	return
}
