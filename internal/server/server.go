package server

import (
	"github.com/sysatom/framework/pkg/config"
	"github.com/sysatom/framework/pkg/flog"

	// File upload handlers
	_ "github.com/sysatom/framework/pkg/media/fs"
	_ "github.com/sysatom/framework/pkg/media/minio"
)

const (
	// Base URL path for serving the streaming API.
	defaultApiPath = "/"
)

func Run() {
	// initialize
	if err := initialize(); err != nil {
		flog.Fatal("initialize %v", err)
	}
	// serve
	if err := listenAndServe(httpApp, config.App.Listen, stopSignal); err != nil {
		flog.Fatal("listenAndServe %v", err)
	}
}
