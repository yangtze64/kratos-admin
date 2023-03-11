package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-admin/app/authorization/internal/biz"
	"kratos-admin/app/authorization/internal/model/sysrole"
	"kratos-admin/pkg/expr"
	"kratos-admin/pkg/global"
	"time"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ro *roleRepo) Create(ctx context.Context, role *biz.Role) (id int32, err error) {
	entry := sysrole.SysRole{
		Name:        role.Name,
		Description: role.Description,
		Operator:    role.Operator,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
	if role.IsEnable {
		entry.IsEnable = 1
	}
	nowTime := time.Now().Unix()
	if role.CreatedAt <= 0 {
		entry.CreatedAt = int32(nowTime)
	}
	if role.UpdatedAt <= 0 {
		entry.UpdatedAt = int32(nowTime)
	}
	err = ro.data.db.WithContext(ctx).Create(&entry).Error
	if err != nil {
		return 0, err
	}
	return entry.Id, nil
}

func (ro *roleRepo) ExistRoleName(ctx context.Context, name string) (exist bool, err error) {
	sub := ro.data.db.WithContext(ctx).Table(sysrole.TableSysRoleName).Select(sysrole.Column.Id.String()).
		Where(sysrole.Column.Name.Eq(), name)
	exist, err = ro.exist(ctx, sub)
	return
}

func (ro *roleRepo) exist(ctx context.Context, sub *gorm.DB) (bool, error) {
	var exist bool
	sub.Where(sysrole.Column.DeletedAt.Eq(), global.ModelNotDeleteAt).Limit(1)
	// "IF(t.id > 0,1,0) as exist"
	st := sysrole.Column.Id.Expr(func(f expr.String) expr.String {
		return "IF(t`" + f + "` > 0, " + expr.Symbol + ", " + expr.Symbol + ") as exist"
	}).String()
	err := ro.data.db.WithContext(ctx).Select(st, 1, 0).Table("(?) as t", sub).Scan(&exist).Error
	if err != nil {
		return false, err
	}
	return exist, nil
}
