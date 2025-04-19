package main

import (
	"github.com/sysatom/framework/internal/server"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
)

// @title						Bussiness API
// @version					1.0
// @description				Bussiness API
// @license.name				Private
// @host						localhost:6060
// @BasePath					/
// @schemes					http
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						X-AccessToken
// @description				access token
func main() {
	fx.New(
		server.Modules,
		// fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		// 	return &fxevent.ZapLogger{Logger: log}
		// }),
	).Run()
}
