package data

import (
	"kratos-admin/app/usercenter/internal/conf"
	"kratos-admin/app/usercenter/internal/data/util"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDefaultDb, NewDefaultRds,
	NewUserRepo)

type (
	DefaultDB  *gorm.DB
	DefaultRDS *redis.Client
	Data       struct {
		db *gorm.DB
		//udb *gorm.DB
		rds *redis.Client
	}
)

func NewDefaultDb(c *conf.Data) (DefaultDB, error) {
	db, err := util.NewGormDB(c.Database.Default)
	if err != nil {
		return nil, err
	}
	return DefaultDB(db), nil
}

//func NewUserDb(c *conf.Data) (UserDB, error) {
//	db, err := util.NewGormDB(c.Database.User)
//	if err != nil {
//		return nil, err
//	}
//	return UserDB(db), nil
//}

func NewDefaultRds(c *conf.Data) DefaultRDS {
	rds := util.NewRedis(c.Redis.Default)
	return DefaultRDS(rds)
}

// NewData .
func NewData(c *conf.Data, db DefaultDB, defRds DefaultRDS, logger log.Logger) (*Data, func(), error) {
	rds := (*redis.Client)(defRds)
	cleanup := func() {
		_ = rds.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: (*gorm.DB)(db), rds: rds}, cleanup, nil
}
