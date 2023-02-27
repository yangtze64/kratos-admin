package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"kratos-admin/app/usercenter/internal/conf"
)

func NewGormDB(c *conf.Data_Mysql) (*gorm.DB, error) {
	db, err := getGormInstance(c, func() *gorm.Config {
		return newGormConf(c)
	})
	if err != nil {
		return nil, err
	}
	conn, _ := db.DB()
	if c.MaxConn > 0 {
		conn.SetMaxOpenConns(int(c.MaxConn))
	}
	if c.MaxIdle > 0 {
		conn.SetMaxIdleConns(int(c.MaxIdle))
	}
	return db, err
}

// 获取gorm实例
func getGormInstance(c *conf.Data_Mysql, fns ...func() *gorm.Config) (*gorm.DB, error) {
	var gc *gorm.Config
	for _, fn := range fns {
		gc = fn()
	}
	db, err := gorm.Open(newGormDial(c), gc)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func newGormConf(c *conf.Data_Mysql) *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	}
}

func newGormDial(c *conf.Data_Mysql) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN:                       c.Dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})
}
