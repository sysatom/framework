package main

import (
	"github.com/sysatom/framework/internal/server"
	// Importing automaxprocs automatically adjusts GOMAXPROCS.
	_ "go.uber.org/automaxprocs"
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
	server.Run()
}
