package controller

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewHelloController, NewValidateController),
	fx.Invoke(BindHelloRoutes, BindValidateRoutes),
)
