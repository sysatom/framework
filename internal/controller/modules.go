package controller

import (
	"github.com/sysatom/framework/internal/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	service.Modules,
	fx.Provide(NewHelloController, NewValidateController, NewUserController),
	fx.Invoke(BindHelloRoutes, BindValidateRoutes, BindAdminRoutes),
)
