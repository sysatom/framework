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

	"github.com/bytedance/sonic"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sysatom/framework/pkg/config"
	"github.com/sysatom/framework/pkg/flog"
	"github.com/sysatom/framework/pkg/utils"
	"github.com/sysatom/framework/version"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

var (
	// swagger
	swagHandler echo.HandlerFunc
	// web app
	httpApp *echo.Echo
)

func Initialize() error {
	var err error

	// init timezone
	if err = initializeTimezone(); err != nil {
		return err
	}
	flog.Info("initialize Timezone ok")

	// init config
	if err = initializeConfig(); err != nil {
		return err
	}
	flog.Info("initialize Config ok")

	// init database
	if err = initializeDatabase(); err != nil {
		return err
	}
	flog.Info("initialize Database ok")

	// init media
	if err = initializeMedia(); err != nil {
		return err
	}
	flog.Info("initialize Media ok")

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
		flog.Fatal("Couldn't get current working directory: %v", err)
	}

	flog.Info("version %s:%s:%s; pid %d; %d process(es)",
		version.Buildtags, executable, version.Buildstamp,
		os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))

	configFile := utils.ToAbsolutePath(curwd, "config.yaml")
	flog.Info("Using config from '%s'", configFile)

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
	flog.Info("API served from root URL path '%s'", config.App.ApiPath)

	// log level
	flog.SetLevel(config.App.Log.Level)

	return nil
}

func NewHTTPServer(lc fx.Lifecycle, logger *zap.Logger) *echo.Echo {
	// Set up HTTP server.
	httpApp = echo.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// setting
			httpApp.HideBanner = true
			httpApp.JSONSerializer = &DefaultJSONSerializer{}
			httpApp.HTTPErrorHandler = func(err error, c echo.Context) {
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
					if httpApp.Debug {
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
					httpApp.Logger.Error(err)
				}
			}

			// middleware
			httpApp.Use(middleware.CORS())
			httpApp.Use(middleware.Recover())
			httpApp.Use(middleware.Decompress())
			httpApp.Use(middleware.Gzip())
			httpApp.Use(middleware.RequestID())
			httpApp.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
				Skipper:      middleware.DefaultSkipper,
				ErrorMessage: "custom timeout error message returns to client",
				OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
					log.Println(c.Path())
				},
				Timeout: 30 * time.Second, // TODO config
			}))
			httpApp.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(200)))) // TODO rate limiter config
			httpApp.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
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
				httpApp.GET("/swagger/*", swagHandler)
			}

			// router
			setupRouter(httpApp)

			go httpApp.Start(config.App.Listen)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpApp.Shutdown(ctx)
		},
	})

	return httpApp
}

func initializeDatabase() error {
	//// init database
	//mysql.Init()
	//store.Init()
	//
	//// Open database
	//err := store.Store.Open(config.App.Store)
	//if err != nil {
	//	return fmt.Errorf("failed to open DB, %w", err)
	//}
	//go func() {
	//	<-stopSignal
	//	err = store.Store.Close()
	//	if err != nil {
	//		flog.Error(err)
	//	}
	//	flog.Debug("Closed database connection(s)")
	//}()
	//
	//// migrate
	//if err := store.Migrate(); err != nil {
	//	return fmt.Errorf("failed to migrate DB, %w", err)
	//}

	return nil
}

func initializeMedia() error {
	// Media
	if config.App.Media != nil {
		if config.App.Media.UseHandler == "" {
			config.App.Media = nil
		} else {
			globals.maxFileUploadSize = config.App.Media.MaxFileUploadSize
			if config.App.Media.Handlers != nil {
				var conf string
				if params := config.App.Media.Handlers[config.App.Media.UseHandler]; params != nil {
					data, err := sonic.Marshal(params)
					if err != nil {
						return fmt.Errorf("failed to marshal media handler, %w", err)
					}
					conf = string(data)
				}
				_, _ = fmt.Println(conf) // FIXME
				//if err := store.UseMediaHandler(config.App.Media.UseHandler, conf); err != nil {
				//	return fmt.Errorf("failed to init media handler, %w", err)
				//}
			}
			if config.App.Media.GcPeriod > 0 && config.App.Media.GcBlockSize > 0 {
				globals.mediaGcPeriod = time.Second * time.Duration(config.App.Media.GcPeriod)
				stopFilesGc := largeFileRunGarbageCollection(globals.mediaGcPeriod, config.App.Media.GcBlockSize)
				go func() {
					// <-stopSignal FIXME
					stopFilesGc <- true
					flog.Info("Stopped files garbage collector")
				}()
			}
		}
	}
	return nil
}
