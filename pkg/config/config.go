package config

import (
	"github.com/sysatom/framework/pkg/utils"
	"github.com/sysatom/framework/version"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// Base URL path for serving the streaming API.
	defaultApiPath = "/"
)

var App Type

// Contentx of the configuration file
type Type struct {
	// HTTP(S) address:port to listen on for websocket and long polling clients. Either a
	// numeric or a canonical name, e.g. ":80" or ":https". Could include a host name, e.g.
	// "localhost:80".
	// Could be blank: if TLS is not configured, will use ":80", otherwise ":443".
	// Can be overridden from the command line, see option --listen.
	Listen string `json:"listen" yaml:"listen" mapstructure:"listen"`
	// Base URL path where the streaming and large file API calls are served, default is '/'.
	// Can be overridden from the command line, see option --api_path.
	ApiPath string `json:"api_path" yaml:"api_path" mapstructure:"api_path"`
	// App Url
	URL string `json:"url" yaml:"url" mapstructure:"url"`

	// Configs for subsystems
	Store StoreType    `json:"store_config" yaml:"store_config" mapstructure:"store_config"`
	Media *mediaConfig `json:"media" yaml:"media" mapstructure:"media"`

	// Redis
	Redis Redis `json:"redis" yaml:"redis" mapstructure:"redis"`

	// Log
	Log Log `json:"log" yaml:"log" mapstructure:"log"`

	// Config for vendors
	Vendors interface{} `json:"vendors" yaml:"vendors" mapstructure:"vendors"`
}

// Large file handler config.
type mediaConfig struct {
	// The name of the handler to use for file uploads.
	UseHandler string `json:"use_handler" yaml:"use_handler" mapstructure:"use_handler"`
	// Maximum allowed size of an uploaded file
	MaxFileUploadSize int64 `json:"max_size" yaml:"max_size" mapstructure:"max_size"`
	// Garbage collection timeout
	GcPeriod int `json:"gc_period" yaml:"gc_period" mapstructure:"gc_period"`
	// Number of entries to delete in one pass
	GcBlockSize int `json:"gc_block_size" yaml:"gc_block_size" mapstructure:"gc_block_size"`
	// Individual handler config params to pass to handlers unchanged.
	Handlers map[string]interface{} `json:"handlers" yaml:"handlers" mapstructure:"handlers"`
}

type StoreType struct {
	// Maximum number of results to return from adapter.
	MaxResults int `json:"max_results" yaml:"max_results" mapstructure:"max_results"`
	// MySQL config
	MySQL MySQL `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
}

type MySQL struct {
	// MySQL DSN
	DSN string `json:"dsn" yaml:"dsn" mapstructure:"dsn"`
	// MySQL max open connections
	MaxOpenConns int `json:"max_open_conns" yaml:"max_open_conns" mapstructure:"max_open_conns"`
	// MySQL max idle connections
	MaxIdleConns int `json:"max_idle_conns" yaml:"max_idle_conns" mapstructure:"max_idle_conns"`
	// MySQL max connection lifetime
	ConnMaxLifetime int `json:"conn_max_lifetime" yaml:"conn_max_lifetime" mapstructure:"conn_max_lifetime"`
	// MySQL max connection lifetime
	SQLTime int `json:"sql_time" yaml:"sql_time" mapstructure:"sql_time"`
}

type Log struct {
	// Log level: debug, info, warn, error, fatal, panic
	Level string `json:"level" yaml:"level" mapstructure:"level"`
}

type Redis struct {
	// Redis host
	Host string `json:"host" yaml:"host" mapstructure:"host"`
	// Redis port
	Port int `json:"port" yaml:"port" mapstructure:"port"`
	// Redis database
	DB int `json:"db" yaml:"db" mapstructure:"db"`
	// Redis password
	Password string `json:"password" yaml:"pass" mapstructure:"password"`
}

func Load(path ...string) {
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatalf("[config] Failed to bind flags: %v", err)
	}
	for _, p := range path {
		viper.AddConfigPath(p)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("[config] Failed to read config file: %v", err)
	}
	err = viper.Unmarshal(&App)
	if err != nil {
		log.Fatalf("[config] Failed to unmarshal config: %v", err)
	}
}

func NewConfig() Type {
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
	Load(".", curwd)

	// Configure root path for serving API calls.
	if App.ApiPath == "" {
		App.ApiPath = defaultApiPath
	} else {
		if !strings.HasPrefix(App.ApiPath, "/") {
			App.ApiPath = "/" + App.ApiPath
		}
		if !strings.HasSuffix(App.ApiPath, "/") {
			App.ApiPath += "/"
		}
	}
	log.Printf("API served from root URL path '%s'\n", App.ApiPath)

	// log level
	// flog.SetLevel(App.Log.Level)

	return App
}
