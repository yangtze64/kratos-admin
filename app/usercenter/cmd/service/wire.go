//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-admin/app/usercenter/internal/biz"
	"kratos-admin/app/usercenter/internal/conf"
	"kratos-admin/app/usercenter/internal/data"
	"kratos-admin/app/usercenter/internal/server"
	"kratos-admin/app/usercenter/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, *conf.JwtAuth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
