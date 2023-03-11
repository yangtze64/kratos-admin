package sysuser

var (
	TableSysUserName = "sys_user"
)

// SysUser mapped from table <sys_user>
// 用户表
type SysUser struct {
	Id        int32   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Uid       string `gorm:"column:uid;type:char(36);comment:UID;NOT NULL" json:"uid"`
	Username  string `gorm:"column:username;type:varchar(100);comment:用户名;NOT NULL" json:"username"`
	Password  string `gorm:"column:password;type:char(32);comment:密码;NOT NULL" json:"password"`
	Realname  string `gorm:"column:realname;type:varchar(100);comment:真实姓名;NOT NULL" json:"realname"`
	Mobile    string `gorm:"column:mobile;type:varchar(15);comment:电话号;NOT NULL" json:"mobile"`
	AreaCode  int32    `gorm:"column:area_code;type:smallint(4);default:86;comment:区号;NOT NULL" json:"area_code"`
	Email     string `gorm:"column:email;type:varchar(255);comment:EMAIL;NOT NULL" json:"email"`
	Weixin    string `gorm:"column:weixin;type:varchar(30);comment:微信号;NOT NULL" json:"weixin"`
	Operator  string `gorm:"column:operator;type:char(36);comment:操作人;NOT NULL" json:"operator"`
	CreatedAt int32    `gorm:"column:created_at;type:int(11);default:0;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt int32    `gorm:"column:updated_at;type:int(11);default:0;comment:修改时间;NOT NULL" json:"updated_at"`
	DeletedAt int32    `gorm:"column:deleted_at;type:int(11);default:0;comment:删除时间;NOT NULL" json:"deleted_at"`
}

func (m *SysUser) TableName() string {
	return TableSysUserName
}
