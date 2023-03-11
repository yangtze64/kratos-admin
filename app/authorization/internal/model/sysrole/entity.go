package sysrole

var (
	TableSysRoleName = "sys_role"
)

// SysRole mapped from table <sys_role>
// 角色表
type SysRole struct {
	Id          int   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name        string `gorm:"column:name;type:varchar(60);comment:角色名;NOT NULL" json:"name"`
	Description string `gorm:"column:description;type:varchar(255);comment:角色描述;NOT NULL" json:"description"`
	IsEnable    int8    `gorm:"column:is_enable;type:tinyint(1);default:0;comment:启用状态 1:启用 0:未启用;NOT NULL" json:"is_enable"`
	Operator    string `gorm:"column:operator;type:char(36);comment:操作人;NOT NULL" json:"operator"`
	CreatedAt   int    `gorm:"column:created_at;type:int(11);default:0;comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt   int    `gorm:"column:updated_at;type:int(11);default:0;comment:修改时间;NOT NULL" json:"updated_at"`
	DeletedAt   int    `gorm:"column:deleted_at;type:int(11);default:0;comment:删除时间;NOT NULL" json:"deleted_at"`
}

func (m *SysRole) TableName() string {
	return TableSysRoleName
}
