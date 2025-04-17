package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sysatom/framework/pkg/types"
	"github.com/sysatom/framework/pkg/types/protocol"
	"net/http"
)

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, protocol.NewSuccessResponse(types.KV{"hello": "world"}))
}
