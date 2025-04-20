package main

import (
	"context"
	"fmt"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/handlers"
	"github.com/DanHerasymenko/GoDelivery/shared/logger"
)

func main() {
	ctx := context.Background()

	// Load config
	cfg, err := config.NewConfigFromEnv()
	if err != nil {
		logger.Fatal(ctx, fmt.Errorf("failed to load config: %w", err))
	}
	logger.Info(ctx, "Config loaded")

	// Create server
	srvr := server.NewServer(cfg)

	// Create handlers
	hdlrs := handlers.NewHandlers(cfg)
	hdlrs.RegisterRoutes(srvr.Router)

	// Run server
	srvr.Run(ctx)

}
