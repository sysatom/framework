package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/samber/oops"
	"github.com/sysatom/framework/pkg/types"
	"github.com/sysatom/framework/pkg/types/protocol"
	"github.com/sysatom/framework/pkg/zlog"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// ShowAccount godoc
//
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	types.FileDef
//	@Router			/hello/{id} [get]
func Hello(c echo.Context) error {
	err := a()
	if err != nil {
		l := zlog.NewZlog()

		l.Info(errors.New("demo").Error())
		l.Info(err.Error(), zap.Error(err))

		return err
	}

	return c.JSON(http.StatusOK, protocol.NewSuccessResponse(types.KV{"hello": "world"}))
}

func d() error {
	return oops.
		Code("iam_missing_permission").
		In("authz").
		Tags("authz").
		Time(time.Now()).
		With("user_id", 1234).
		With("permission", "post.create").
		Hint("Runbook: https://doc.acme.org/doc/abcd.md").
		User("user-123", "firstname", "john", "lastname", "doe").
		Errorf("permission denied")
}

func c() error {
	return d()
}

func b() error {
	// add more context
	return oops.
		In("iam").
		Tags("iam").
		Trace("e76031ee-a0c4-4a80-88cb-17086fdd19c0").
		With("hello", "world").
		Wrapf(c(), "something failed")
}

func a() error {
	return b()
}
