package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"time"

	"github.com/sysatom/framework/pkg/cache"
	"github.com/sysatom/framework/pkg/flog"
)

func listenAndServe(app *echo.Echo, addr string, stop <-chan bool) error {
	globals.shuttingDown = false

	httpdone := make(chan bool)

	go func() {
		err := app.Start(addr)
		if err != nil {
			flog.Error(err)
		}
		httpdone <- true
	}()

	// Wait for either a termination signal or an error
Loop:
	for {
		select {
		case <-stop:
			// Flip the flag that we are terminating and close the Accept-ing socket, so no new connections are possible.
			globals.shuttingDown = true
			// Give server 2 seconds to shut down.
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			if err := app.Shutdown(ctx); err != nil {
				// failure/timeout shutting down the server gracefully
				flog.Error(err)
			}

			cancel()

			// Shutdown Extra
			//globals.taskQueue.Shutdown()
			//globals.manager.Shutdown()
			//globals.cronTaskManager.Shutdown()
			//for _, ruleset := range globals.cronRuleset {
			//	ruleset.Shutdown()
			//}
			cache.Shutdown()

			break Loop
		case <-httpdone:
			break Loop
		}
	}
	return nil
}
