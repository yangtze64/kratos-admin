package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-admin/utils"
	"kratos-admin/utils/global"
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

type UserRepo interface {
	Create(ctx context.Context, user *User) (id int, err error)
	Update(ctx context.Context, uid string, user *User) error
	List(ctx context.Context, user *User) (list []*User, total int64, err error)
	FindByUid(ctx context.Context, uid string) (*User, error)
	ExistUser(ctx context.Context, uid string) (bool, error)
	ExistUsername(ctx context.Context, username string, excludeUids ...string) (bool, error)
	ExistMobile(ctx context.Context, mobile string, areaCode int32, excludeUids ...string) (bool, error)
	ExistEmail(ctx context.Context, email string, excludeUids ...string) (bool, error)
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
