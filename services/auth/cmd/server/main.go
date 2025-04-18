package main

import (
	"context"
	"fmt"
	"github.com/DanHerasymenko/GoDelivery/shared/config"
	"github.com/DanHerasymenko/GoDelivery/shared/logger"
	"log/slog"
)

func main() {
	ctx := context.Background()

	// Load config
	_, err := config.NewConfigFromEnv()
	if err != nil {
		logger.Fatal(ctx, fmt.Errorf("failed to load config: %w", err))
	}
	logger.Info(ctx, "Config loaded")

	ctx = logger.WithAttr(ctx, slog.String("service", "auth-service"))
	logger.Info(ctx, "Service started")

}
