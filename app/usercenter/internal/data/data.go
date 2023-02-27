package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
	"kratos-admin/app/usercenter/internal/conf"
	"kratos-admin/app/usercenter/internal/data/util"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewPassportRepo)

// Data .
type Data struct {
	db  *gorm.DB
	udb *gorm.DB
	rds *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := util.NewGormDB(c.Database.Default)
	if err != nil {
		return nil, nil, err
	}
	udb, err := util.NewGormDB(c.Database.User)
	if err != nil {
		return nil, nil, err
	}
	rds := util.NewRedis(c.Redis.Default)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		_ = rds.Close()
	}
	return &Data{db: db, udb: udb, rds: rds}, cleanup, nil
}
