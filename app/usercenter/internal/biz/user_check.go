package biz

import (
	"context"
	"kratos-admin/pkg/errx"
)

type UserCheckOption func(ctx context.Context, user *User) error

func UserCheckChain(ctx context.Context, user *User, opts ...UserCheckOption) error {
	if len(opts) > 0 {
		for _, opt := range opts {
			if err := opt(ctx, user); err != nil {
				return err
			}
		}
	}
	return nil
}

func WithNotExistUser(repo UserRepo) UserCheckOption {
	return func(ctx context.Context, user *User) error {
		ok, err := repo.ExistUser(ctx, user.Uid)
		if err != nil {
			return err
		}
		if !ok {
			return errx.New(errx.UserNotFound)
		}
		return nil
	}
}

func WithExistUsername(repo UserRepo, excludeUids ...string) UserCheckOption {
	return func(ctx context.Context, user *User) error {
		ok, err := repo.ExistUsername(ctx, user.Username, excludeUids...)
		if err != nil {
			return err
		}
		if ok {
			return errx.New(errx.UserNameExist)
		}
		return nil
	}
}

func WithExistMobile(repo UserRepo, excludeUids ...string) UserCheckOption {
	return func(ctx context.Context, user *User) error {
		ok, err := repo.ExistMobile(ctx, user.Mobile, user.AreaCode, excludeUids...)
		if err != nil {
			return err
		}
		if ok {
			return errx.New(errx.UserMobileExist)
		}
		return nil
	}
}

func WithExistEmail(repo UserRepo, excludeUids ...string) UserCheckOption {
	return func(ctx context.Context, user *User) error {
		ok, err := repo.ExistEmail(ctx, user.Email, excludeUids...)
		if err != nil {
			return err
		}
		if ok {
			return errx.New(errx.UserEmailExist)
		}
		return nil
	}
}
