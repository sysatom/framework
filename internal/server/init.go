package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/pflag"
	"github.com/sysatom/framework/pkg/cache"
	"github.com/sysatom/framework/pkg/config"
	"github.com/sysatom/framework/pkg/event"
	"github.com/sysatom/framework/pkg/flog"
	"github.com/sysatom/framework/pkg/types"
	"github.com/sysatom/framework/pkg/utils"
	"github.com/sysatom/framework/version"
	"github.com/ziflex/lecho/v3"
	"golang.org/x/time/rate"
)

var (
	// stop signal
	stopSignal <-chan bool
	// swagger
	swagHandler echo.HandlerFunc
	// web app
	httpApp *echo.Echo
	// flag variables
	appFlag struct {
		configFile *string
		listenOn   *string
		apiPath    *string
		tlsEnabled *bool
		pprofFile  *string
		pprofUrl   *string
	}
)

func initialize() error {
	var err error

	// init log
	if err = initializeLog(); err != nil {
		return err
	}
	flog.Info("initialize Log ok")

	// init timezone
	if err = initializeTimezone(); err != nil {
		return err
	}
	flog.Info("initialize Timezone ok")

	// init flag
	if err = initializeFlag(); err != nil {
		return err
	}
	flog.Info("initialize Flag ok")

	// init config
	if err = initializeConfig(); err != nil {
		return err
	}
	flog.Info("initialize Config ok")

	// init http
	if err = initializeHttp(); err != nil {
		return err
	}
	flog.Info("initialize Http ok")

	// init pprof
	if err = initializePprof(); err != nil {
		return err
	}
	flog.Info("initialize Pprof ok")

	// init cache
	if err = initializeCache(); err != nil {
		return err
	}
	flog.Info("initialize Cache ok")

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

	// init signal
	if err = initializeSignal(); err != nil {
		return err
	}
	flog.Info("initialize Signal ok")

	// init event
	if err = initializeEvent(); err != nil {
		return err
	}
	flog.Info("initialize Event ok")

	return nil
}

func initializeLog() error {
	flog.Init(false)
	return nil
}

func initializeTimezone() error {
	_, err := time.LoadLocation("Local")
	if err != nil {
		return fmt.Errorf("load time location error, %w", err)
	}
	return nil
}

func initializeFlag() error {
	appFlag.configFile = pflag.String("config", "flowbot.yaml", "Path to config file.")
	appFlag.listenOn = pflag.String("listen", "", "Override address and port to listen on for HTTP(S) clients.")
	appFlag.apiPath = pflag.String("api_path", "", "Override the base URL path where API is served.")
	appFlag.tlsEnabled = pflag.Bool("tls_enabled", false, "Override config value for enabling TLS.")
	appFlag.pprofFile = pflag.String("pprof", "", "File name to save profiling info to. Disabled if not set.")
	appFlag.pprofUrl = pflag.String("pprof_url", "", "Debugging only! URL path for exposing profiling info. Disabled if not set.")
	pflag.Parse()
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

	*appFlag.configFile = utils.ToAbsolutePath(curwd, *appFlag.configFile)
	flog.Info("Using config from '%s'", *appFlag.configFile)

	// Load config
	config.Load(".", curwd)

	if *appFlag.listenOn != "" {
		config.App.Listen = *appFlag.listenOn
	}

	// Configure root path for serving API calls.
	if *appFlag.apiPath != "" {
		config.App.ApiPath = *appFlag.apiPath
	}
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

// DefaultJSONSerializer implements JSON encoding using encoding/json.
type DefaultJSONSerializer struct{}

// Serialize converts an interface into a json and writes it to the response.
// You can optionally use the indent parameter to produce pretty JSONs.
func (d DefaultJSONSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := sonic.ConfigDefault.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

// Deserialize reads a JSON from a request body and converts it into an interface.
func (d DefaultJSONSerializer) Deserialize(c echo.Context, i interface{}) error {
	err := sonic.ConfigDefault.NewDecoder(c.Request().Body).Decode(i)
	if ute, ok := err.(*json.UnmarshalTypeError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
	} else if se, ok := err.(*json.SyntaxError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
	}
	return err
}

func initializeHttp() error {
	// Set up HTTP server.
	httpApp = echo.New()

	// setting
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
	httpApp.Logger = lecho.From(flog.GetLogger())

	// middleware
	//httpApp.Use(middleware.Logger())
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
	httpApp.Use(lecho.Middleware(lecho.Config{
		Logger: lecho.From(flog.GetLogger()),
	}))

	// hook
	//httpApp.Hooks().OnRoute(func(r fiber.Route) error {
	//	if r.Method == http.MethodHead {
	//		return nil
	//	}
	//	flog.Info("[route] %+7s %s", r.Method, r.Path)
	//	return nil
	//})

	// swagger
	if swagHandler != nil {
		httpApp.GET("/swagger/*", swagHandler)
	}

	// mux router
	setupMux(httpApp)

	return nil
}

func initializePprof() error {
	// Initialize serving debug profiles (optional).
	//pprofs.ServePprof(httpApp, *appFlag.pprofUrl)

	if *appFlag.pprofFile != "" {
		curwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current working directory, %w", err)
		}
		*appFlag.pprofFile = utils.ToAbsolutePath(curwd, *appFlag.pprofFile)

		cpuf, err := os.Create(*appFlag.pprofFile + ".cpu")
		if err != nil {
			flog.Fatal("Failed to create CPU pprof file: %v", err)
		}
		defer func() {
			_ = cpuf.Close()
		}()

		memf, err := os.Create(*appFlag.pprofFile + ".mem")
		if err != nil {
			flog.Fatal("Failed to create Mem pprof file: %v", err)
		}
		defer func() {
			_ = memf.Close()
		}()

		_ = pprof.StartCPUProfile(cpuf)
		defer pprof.StopCPUProfile()
		defer func() {
			_ = pprof.WriteHeapProfile(memf)
		}()

		flog.Info("Profiling info saved to '%s.(cpu|mem)'", *appFlag.pprofFile)
	}
	return nil
}

func initializeCache() error {
	// init cache
	return cache.InitCache()
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
					<-stopSignal
					stopFilesGc <- true
					flog.Info("Stopped files garbage collector")
				}()
			}
		}
	}
	return nil
}

func initializeSignal() error {
	stopSignal = utils.SignalHandler()
	return nil
}

// init event
func initializeEvent() error {
	router, err := event.NewRouter()
	if err != nil {
		return err
	}

	subscriber, err := event.NewSubscriber()
	if err != nil {
		return err
	}

	router.AddNoPublisherHandler(
		"onInstructPushEventHandler",
		types.InstructPushEvent,
		subscriber,
		onInstructPushEventHandler,
	)

	go func() {
		if err = router.Run(context.Background()); err != nil {
			flog.Error(err)
		}
	}()

	return nil
}
