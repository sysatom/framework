package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

var (
	// swagger
	swagHandler echo.HandlerFunc
)

func initialize() error {
	var err error

	// init timezone
	if err = initializeTimezone(); err != nil {
		return err
	}
	log.Println("initialize Timezone ok")

	return nil
}

func initializeTimezone() error {
	_, err := time.LoadLocation("Local")
	if err != nil {
		return fmt.Errorf("load time location error, %w", err)
	}
	return nil
}
