package main

import (
	"GoDelivery/shared/config"
	"GoDelivery/shared/logger"
	"context"
	"fmt"
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
