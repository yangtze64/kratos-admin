package model

import (
	"database/sql"
	"time"
)

var TableSysUserName = "sys_user"

// SysUser mapped from table <sys_user>
type SysUser struct {
	ID       int          `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UID      string       `gorm:"column:uid;type:char(36);comment:UID;NOT NULL" json:"uid"`
	Username string       `gorm:"column:username;type:varchar(100);comment:用户名;NOT NULL" json:"username"`
	Password string       `gorm:"column:password;type:char(32);comment:密码;NOT NULL" json:"password"`
	Realname string       `gorm:"column:realname;type:varchar(100);comment:真实姓名;NOT NULL" json:"realname"`
	Mobile   string       `gorm:"column:mobile;type:varchar(15);comment:电话号;NOT NULL" json:"mobile"`
	AreaCode int32        `gorm:"column:area_code;type:smallint(4) unsigned;default:86;comment:区号;NOT NULL" json:"area_code"`
	Email    string       `gorm:"column:email;type:varchar(255);comment:EMAIL;NOT NULL" json:"email"`
	Weixin   string       `gorm:"column:weixin;type:varchar(30);comment:微信号;NOT NULL" json:"weixin"`
	Unionid  string       `gorm:"column:unionid;type:varchar(64);comment:微信平台UnionID;NOT NULL" json:"unionid"`
	CreateAt time.Time    `gorm:"column:create_at;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间;NOT NULL" json:"create_at"`
	UpdateAt time.Time    `gorm:"column:update_at;type:datetime;default:CURRENT_TIMESTAMP;comment:修改时间;NOT NULL" json:"update_at"`
	Operator string       `gorm:"column:operator;type:char(36);comment:操作人;NOT NULL" json:"operator"`
	IsDelete int8         `gorm:"column:is_delete;type:tinyint(1);default:0;comment:是否删除;NOT NULL" json:"is_delete"`
	DeleteAt sql.NullTime `gorm:"column:delete_at;type:datetime;comment:删除时间" json:"delete_at"`
}

func (m *SysUser) TableName() string {
	return TableSysUserName
}
