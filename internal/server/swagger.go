//go:build swagger

package server

import (
	"github.com/gofiber/swagger"
	_ "github.com/sysatom/framework/docs"
)

func init() {
	swagHandler = swagger.HandlerDefault
}
