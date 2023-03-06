package data

import (
	"gorm.io/gorm"
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/app/usercenter/internal/model/sysuser"
	"kratos-admin/utils/global"
	"time"
)

type UserCondOption func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error)

func UserCondChain(db *gorm.DB, user *biz.User, opts ...UserCondOption) (tx *gorm.DB, err error) {
	tx = db
	for _, opt := range opts {
		if tx, err = opt(tx, user); err != nil {
			return tx, err
		}
	}
	return tx, nil
}

func WithIdWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.Id > 0 {
			tx.Where(sysuser.Column.Id.Eq(), user.Id)
		} else {
			if n := len(user.Ids); n > 0 {
				if n == 1 {
					tx.Where(sysuser.Column.Id.Eq(), user.Ids[0])
				} else {
					tx.Where(sysuser.Column.Id.In(), user.Ids)
				}
			}
		}
		return
	}
}
func WithUidWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.Uid != "" {
			tx.Where(sysuser.Column.Uid.Eq(), user.Uid)
		} else {
			if n := len(user.Uids); n > 0 {
				if n == 1 {
					tx.Where(sysuser.Column.Uid.Eq(), user.Uids[0])
				} else {
					tx.Where(sysuser.Column.Uid.In(), user.Uids)
				}
			}
		}
		return
	}
}
func WithMobileWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.Mobile != "" {
			tx.Where(sysuser.Column.Mobile.Eq(), user.Mobile)
		}
		if user.AreaCode > 0 {
			tx.Where(sysuser.Column.AreaCode.Eq(), user.AreaCode)
		}
		return
	}
}
func WithUsernameWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.Username != "" {
			tx.Where(sysuser.Column.Username.Eq(), user.Username)
		}
		return
	}
}
func WithUsernameFuzzyWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.UsernameFuzzy != "" {
			tx.Where(sysuser.Column.Username.Like(), user.UsernameFuzzy+"%")
		}
		return
	}
}
func WithEmailWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.Email != "" {
			tx.Where(sysuser.Column.Email.Eq(), user.Email)
		}
		return
	}
}
func WithEmailFuzzyWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.EmailFuzzy != "" {
			tx.Where(sysuser.Column.Email.Like(), user.EmailFuzzy+"%")
		}
		return
	}
}
func WithRealnameWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.Realname != "" {
			tx.Where(sysuser.Column.Realname.Eq(), user.Realname)
		}
		return
	}
}
func WithRealnameFuzzyWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.RealnameFuzzy != "" {
			tx.Where(sysuser.Column.Realname.Like(), user.RealnameFuzzy+"%")
		}
		return
	}
}
func WithWeixinWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.Weixin != "" {
			tx.Where(sysuser.Column.Weixin.Eq(), user.Weixin)
		}
		return
	}
}
func WithOperatorWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.Operator != "" {
			tx.Where(sysuser.Column.Operator.Eq(), user.Operator)
		} else {
			if n := len(user.Operators); n > 0 {
				if n == 1 {
					tx.Where(sysuser.Column.Operator.Eq(), user.Operators[0])
				} else {
					tx.Where(sysuser.Column.Operator.In(), user.Operators)
				}
			}
		}
		return
	}
}
func WithCreatedAtWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.CreatedAt > 0 {
			tx.Where(sysuser.Column.CreatedAt.Eq(), user.CreatedAt)
		} else {
			if user.CreatedDateStart != "" {
				timeS, err := time.ParseInLocation(global.TimeFormat, user.CreatedDateStart, time.Local)
				if err != nil {
					return tx, err
				}
				tx.Where(sysuser.Column.CreatedAt.Gte(),timeS.Unix())
			}
			if user.CreatedDateEnd != "" {
				timeE, err := time.ParseInLocation(global.TimeFormat, user.CreatedDateEnd, time.Local)
				if err != nil {
					return tx, err
				}
				tx.Where(sysuser.Column.CreatedAt.Lte(),timeE.Unix())
			}
		}
		return
	}
}
func WithUpdatedAtWhere() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.UpdatedAt > 0 {
			tx.Where(sysuser.Column.UpdatedAt.Eq(), user.UpdatedAt)
		} else {
			if user.UpdatedDateStart != "" {
				timeS, err := time.ParseInLocation(global.TimeFormat, user.UpdatedDateStart, time.Local)
				if err != nil {
					return tx, err
				}
				tx.Where(sysuser.Column.UpdatedAt.Gte(),timeS.Unix())
			}
			if user.UpdatedDateEnd != "" {
				timeE, err := time.ParseInLocation(global.TimeFormat, user.UpdatedDateEnd, time.Local)
				if err != nil {
					return tx, err
				}
				tx.Where(sysuser.Column.UpdatedAt.Lte(),timeE.Unix())
			}
		}
		return
	}
}
func WithLimit() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.Limit > 0 {
			tx.Limit(user.Limit)
		}
		return
	}
}
func WithPager() UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if user.Limit > 0 && user.Page > 0{
			tx.Limit(user.Limit).Offset((user.Page - 1) * user.Limit)
		}
		return
	}
}
func WithOrder(field string, asc bool) UserCondOption {
	return func(db *gorm.DB, user *biz.User) (tx *gorm.DB, err error) {
		tx = db
		if asc {
			tx.Order(field+" ASC")
		}else{
			tx.Order(field+" DESC")
		}
		return
	}
}


