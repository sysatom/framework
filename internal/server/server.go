package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sysatom/framework/pkg/config"
	"github.com/sysatom/framework/pkg/utils"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"time"
)

func NewHTTPServer(config.Type) *echo.Echo {
	return echo.New()
}

func RegisterHooks(lc fx.Lifecycle, httpServer *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// initialize
			if err := initialize(); err != nil {
				return fmt.Errorf("failed to initialize: %w", err)
			}

			// setting
			httpServer.HideBanner = true
			httpServer.JSONSerializer = &utils.DefaultJSONSerializer{}
			httpServer.Validator = &CustomValidator{validator.New()}
			httpServer.HTTPErrorHandler = func(err error, c echo.Context) {
				if c.Response().Committed {
					return
				}

				he, ok := err.(*echo.HTTPError)
				if ok {
					if he.Internal != nil {
						if herr, ok := he.Internal.(*echo.HTTPError); ok {
							he = herr
						}
					}
				} else {
					he = &echo.HTTPError{
						Code:    http.StatusInternalServerError,
						Message: http.StatusText(http.StatusInternalServerError),
					}
				}

				// Issue #1426
				code := he.Code
				message := he.Message

				switch m := he.Message.(type) {
				case string:
					if httpServer.Debug {
						message = echo.Map{"message": m, "error": err.Error()}
					} else {
						message = echo.Map{"message": m}
					}
				case json.Marshaler:
					// do nothing - this type knows how to format itself to JSON
				case error:
					message = echo.Map{"message": m.Error()}
				}

				// Send response
				if c.Request().Method == http.MethodHead { // Issue #608
					err = c.NoContent(he.Code)
				} else {
					err = c.JSON(code, message)
				}
				if err != nil {
					httpServer.Logger.Error(err)
				}
			}

			// middleware
			httpServer.Use(middleware.CORS())
			httpServer.Use(middleware.Recover())
			httpServer.Use(middleware.Decompress())
			httpServer.Use(middleware.Gzip())
			httpServer.Use(middleware.RequestID())
			httpServer.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
				Skipper:      middleware.DefaultSkipper,
				ErrorMessage: "custom timeout error message returns to client",
				OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
					log.Println(c.Path())
				},
				Timeout: 30 * time.Second, // TODO config
			}))
			httpServer.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(200)))) // TODO rate limiter config
			httpServer.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
				LogURI:    true,
				LogStatus: true,
				Skipper: func(c echo.Context) bool {
					// Skip health check endpoint
					return c.Request().URL.Path == "/health"
				},
				LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
					c.Logger().Info("request",
						zap.String("URI", v.URI),
						zap.Int("status", v.Status),
					)

					return nil
				},
			}))

			// swagger
			if swagHandler != nil {
				httpServer.GET("/swagger/*", swagHandler)
			}

			// router
			//setupRouter(httpServer)

			go func() {
				err := httpServer.Start(config.App.Listen)
				if err != nil {
					httpServer.Logger.Panic(err.Error())
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})
}

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
