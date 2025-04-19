package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sysatom/framework/internal/service"
	"github.com/sysatom/framework/pkg/types"
	"github.com/sysatom/framework/pkg/types/protocol"
	"go.uber.org/zap"
	"net/http"
)

type UserController struct {
	userService *service.UserService
	logger      *zap.Logger
}

func NewUserController(logger *zap.Logger, userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
		logger:      logger,
	}
}

// Login godoc
//
//	@Summary		User login
//	@Description	User login with username and password
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			body	body		types.UserLoginRequest	true	"Login credentials"
//	@Success		200		{object}	protocol.Response{data=types.UserLoginResponse}
//	@Router			/user/login [post]
func (controller UserController) Login(c echo.Context) error {
	var user types.UserLoginRequest
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(&user); err != nil {
		return err
	}

	// login
	id, token, err := controller.userService.Login(c.Request().Context(), user.Username, user.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, protocol.NewSuccessResponse(types.UserLoginResponse{
		ID:    types.Uint64ID(id),
		Token: token,
	}))
}

// Info godoc
//
//	@Summary		User info
//	@Description	User info
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	protocol.Response{data=types.KV}
//	@Router			/user/info [get]
func (controller UserController) Info(c echo.Context) error {
	return c.JSON(http.StatusOK, protocol.NewSuccessResponse(types.KV{}))
}
