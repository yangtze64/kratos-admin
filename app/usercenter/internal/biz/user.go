package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-admin/utils"
	"time"
)

type User struct {
	Id       int
	Uid      string
	Username string
	Realname string
	Password string
	Mobile   string
	AreaCode int32
	Email    string
	Weixin   string
	Unionid  string
	CreateAt time.Time
	UpdateAt time.Time
	Operator string
}

type UserRepo interface {
	Create(ctx context.Context, user *User) (id int, err error)
	Update(ctx context.Context, uid string) (bool, error)
	Delete(ctx context.Context, uid string) (bool, error)
	List(ctx context.Context) ([]*User, error)
	FindByUid(ctx context.Context, uid string) (*User, error)
	ExistUsername(ctx context.Context, username string) (bool, error)
	ExistMobile(ctx context.Context, mobile string, areaCode int32) (bool, error)
	ExistEmail(ctx context.Context, email string) (bool, error)
	ExistUnionId(ctx context.Context, unionid string) (bool, error)
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
	if user.Unionid != "" {
		checkOptions = append(checkOptions, WithExistUnionId(u.repo))
	}
	if user.AreaCode == 0 {
		user.AreaCode = 86
	}
	err := UserCheckChain(ctx, user, checkOptions...)
	if err != nil {
		return nil, err
	}
	user.Uid = utils.NewUuid()
	user.Password = utils.GenPasswd(user.Password)
	nowTime := time.Now()
	if user.CreateAt.IsZero() {
		user.CreateAt = nowTime
	}
	if user.UpdateAt.IsZero() {
		user.CreateAt = nowTime
	}
	id, err := u.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = id
	return user, nil
}
