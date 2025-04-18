package types

import (
	"github.com/lithammer/shortuuid/v4"
	"github.com/sysatom/framework/pkg/config"
)

func Id() string {
	return shortuuid.New()
}

func AppUrl() string {
	return config.App.URL
}
