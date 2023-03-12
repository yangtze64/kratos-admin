package data

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	etcdclient "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
	"kratos-admin/app/authorization/internal/conf"
	"kratos-admin/app/authorization/internal/data/util"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewRegistrar, NewDefaultDb, NewRds, NewRoleRepo)

type (
	DefaultDB *gorm.DB
	Data      struct {
		db  *gorm.DB
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

func NewRds(c *conf.Data) (*redis.Client, error) {
	rds, err := util.NewRedis(c.Redis)
	if err != nil {
		return nil, err
	}
	return rds, err
}

func NewRegistrar(conf *conf.Registry) (registry.Registrar, error) {
	point := conf.Etcd.Address
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{point},
	})
	if err != nil {
		return nil, err
	}
	r := etcd.New(client)
	return r, nil
}

// NewData .
func NewData(c *conf.Data, db DefaultDB, rds *redis.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		_ = rds.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: (*gorm.DB)(db), rds: rds}, cleanup, nil
}
