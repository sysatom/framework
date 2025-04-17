package main

import (
	"github.com/sysatom/framework/internal/server"
	// Importing automaxprocs automatically adjusts GOMAXPROCS.
	_ "go.uber.org/automaxprocs"
)

// @title						Flowbot API
// @version					1.0
// @description				Flowbot Chatbot API
// @license.name				GPL 3.0
// @license.url				https://github.com/sysatom/framework/blob/master/LICENSE
// @host						localhost:6060
// @BasePath					/service
// @schemes					http
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						X-AccessToken
// @description				access token
func main() {
	server.Run()
}
