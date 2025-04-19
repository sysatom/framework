package service

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(NewHelloService),
)
