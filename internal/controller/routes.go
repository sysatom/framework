package controller

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sysatom/framework/pkg/config"
)

func BindHelloRoutes(e *echo.Echo, c *HelloController) {
	e.GET("/hello", c.Hello)
	e.GET("/ent", c.Ent)
}

func BindValidateRoutes(e *echo.Echo, c *ValidateController) {
	e.GET("/validate", c.Validate)

	g := e.Group("/auth")
	g.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: "secret",
	})) // TODO jwt config
	g.GET("/some", c.Validate)
}

func BindAdminRoutes(_ config.Type, e *echo.Echo, c *UserController) {
	g := e.Group("/user")
	g.POST("/login", c.Login)

	//g.Use(echojwt.WithConfig(echojwt.Config{
	//	SigningKey: config.App.JWTSecret,
	//}))
	g.GET("/info", c.Info)
}
