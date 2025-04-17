package server

import (
	"github.com/labstack/echo/v4"
	"github.com/sysatom/framework/internal/controller"
)

// router
func setupMux(a *echo.Echo) {
	a.GET("/hello", controller.Hello)
}
