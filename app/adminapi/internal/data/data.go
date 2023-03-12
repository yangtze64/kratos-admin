package data

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	etcdclient "go.etcd.io/etcd/client/v3"
	authorizationv1 "kratos-admin/api/authorization/service/v1"
	userCenterv1 "kratos-admin/api/usercenter/service/v1"
	"kratos-admin/app/adminapi/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRegistrar,
	NewDiscovery,
	NewUserCenterClient,
	NewAuthorizationClient,
	NewUserCenterRepo,
	NewAuthorizationRepo,
)

type Data struct {
	UserCenterClient    userCenterv1.UserCenterClient
	AuthorizationClient authorizationv1.AuthorizationClient
}

func NewRegistrar(conf *conf.Registry) registry.Registrar{
	point := conf.Etcd.Address
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{point},
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	point := conf.Etcd.Address
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{point},
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}

// NewData .
func NewData(c *conf.Data, userCenterClient userCenterv1.UserCenterClient, authorizationClient authorizationv1.AuthorizationClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{UserCenterClient: userCenterClient, AuthorizationClient: authorizationClient}, cleanup, nil
}
