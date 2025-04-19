package repository

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(NewMySQLClient, NewRedisClient),
)
