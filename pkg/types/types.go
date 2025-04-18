package types

import (
	"fmt"
	"github.com/sony/sonyflake"
	"github.com/sysatom/framework/pkg/config"
	"time"
)

func UniqueId() (uint64, error) {
	startTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	s, err := sonyflake.New(sonyflake.Settings{
		StartTime: startTime,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to create sonyflake, %w", err)
	}
	return s.NextID()
}

func AppUrl() string {
	return config.App.URL
}
