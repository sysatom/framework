package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sysatom/framework/internal/server"
	"github.com/sysatom/framework/pkg/rdb"
	"github.com/sysatom/framework/pkg/zlog"

	// Importing automaxprocs automatically adjusts GOMAXPROCS.
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
	// initialize
	if err := server.Initialize(); err != nil {
		log.Fatalf("initialize %v", err)
	}
	// serve
	fx.New(
		fx.Provide(
			server.NewHTTPServer,
			zlog.NewZlog,
			rdb.NewRedisClient,
		),
		// fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		// 	return &fxevent.ZapLogger{Logger: log}
		// }),
		fx.Invoke(func(*echo.Echo) {}),
	).Run()
}
