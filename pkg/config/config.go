package config

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var App configType

// Contentx of the configuration file
type configType struct {
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
	// DB adapter name to use. Should be one of those specified in `Adapters`.
	UseAdapter string `json:"use_adapter" yaml:"use_adapter" mapstructure:"use_adapter"`
	// Configurations for individual adapters.
	Adapters map[string]interface{} `json:"adapters" yaml:"adapters" mapstructure:"adapters"`
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
