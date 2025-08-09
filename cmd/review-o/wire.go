//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/yygqzhu/review-o/internal/biz"
	"github.com/yygqzhu/review-o/internal/conf"
	"github.com/yygqzhu/review-o/internal/data"
	"github.com/yygqzhu/review-o/internal/server"
	"github.com/yygqzhu/review-o/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
