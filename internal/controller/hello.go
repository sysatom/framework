package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sysatom/framework/pkg/types"
	"github.com/sysatom/framework/pkg/types/protocol"
	"net/http"
)

// ShowAccount godoc
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	types.WorkflowTask
//	@Router			/hello/{id} [get]
func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, protocol.NewSuccessResponse(types.KV{"hello": "world"}))
}
