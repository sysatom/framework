package server

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sysatom/framework/internal/controller"
)

// router
func setupRouter(a *echo.Echo) {
	a.GET("/hello", controller.Hello)
	a.GET("/validate", controller.Validate)

	g := a.Group("/auth")
	g.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: "secret",
	})) // TODO jwt config
	g.GET("/some", controller.Hello)
}
