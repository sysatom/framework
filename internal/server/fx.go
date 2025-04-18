package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sysatom/framework/pkg/config"
	"github.com/sysatom/framework/pkg/utils"
	"github.com/sysatom/framework/version"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

var (
	// swagger
	swagHandler echo.HandlerFunc
)

func Initialize() error {
	var err error

	// init timezone
	if err = initializeTimezone(); err != nil {
		return err
	}
	log.Println("initialize Timezone ok")

	// init config
	if err = initializeConfig(); err != nil {
		return err
	}
	log.Println("initialize Config ok")

	return nil
}

func initializeTimezone() error {
	_, err := time.LoadLocation("Local")
	if err != nil {
		return fmt.Errorf("load time location error, %w", err)
	}
	return nil
}

func initializeConfig() error {
	executable, _ := os.Executable()

	curwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Couldn't get current working directory: %v", err)
	}

	log.Printf("version %s:%s:%s; pid %d; %d process(es)\n",
		version.Buildtags, executable, version.Buildstamp,
		os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))

	configFile := utils.ToAbsolutePath(curwd, "config.yaml")
	log.Printf("Using config from '%s'\n", configFile)

	// Load config
	config.Load(".", curwd)

	// Configure root path for serving API calls.
	if config.App.ApiPath == "" {
		config.App.ApiPath = defaultApiPath
	} else {
		if !strings.HasPrefix(config.App.ApiPath, "/") {
			config.App.ApiPath = "/" + config.App.ApiPath
		}
		if !strings.HasSuffix(config.App.ApiPath, "/") {
			config.App.ApiPath += "/"
		}
	}
	log.Printf("API served from root URL path '%s'\n", config.App.ApiPath)

	// log level
	// flog.SetLevel(config.App.Log.Level)

	return nil
}

func NewHTTPServer(lc fx.Lifecycle, logger *zap.Logger) *echo.Echo {
	// Set up HTTP server.
	httpServer := echo.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// setting
			httpServer.HideBanner = true
			httpServer.JSONSerializer = &DefaultJSONSerializer{}
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
					logger.Info("request",
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
			setupRouter(httpServer)

			go func() {
				err := httpServer.Start(config.App.Listen)
				if err != nil {
					logger.Panic(err.Error())
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})

	return httpServer
}
