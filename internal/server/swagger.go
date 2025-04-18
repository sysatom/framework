//go:build swagger

package server

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/sysatom/framework/docs"
)

func init() {
	swagHandler = echoSwagger.WrapHandler
}
