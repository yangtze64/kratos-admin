package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	Uid       string
	Username  string `validate:"required"`
	Realname  string
	Password  string
	Mobile    string
	AreaCode  int32
	Email     string
	Weixin    string
	Unionid   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Operator  string
}

type UserRepo interface {
	Create(ctx context.Context, user *User) (uid int64, err error)
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
	fmt.Println(user)
	return user, nil
}
