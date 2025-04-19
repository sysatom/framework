package server

import (
	"github.com/sysatom/framework/internal/controller"
	"github.com/sysatom/framework/internal/repository"
	"github.com/sysatom/framework/pkg/config"
	"github.com/sysatom/framework/pkg/zlog"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	controller.Modules,
	repository.Modules,
	fx.Provide(NewHTTPServer, config.NewConfig, zlog.NewZlog),
	fx.Invoke(RegisterHooks),
)
