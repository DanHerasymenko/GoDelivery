// @title           GoDelivery Auth Service
// @version         1.0
// @description     This is a sample server for GoDelivery Auth Service.

package main

import (
	"context"
	"fmt"
	_ "github.com/DanHerasymenko/GoDelivery/services/auth-service/cmd/server/docs"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/handlers"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/server/middleware"
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

	// Create clients
	clnts, err := config.NewClients(ctx, cfg)
	if err != nil {
		logger.Fatal(ctx, fmt.Errorf("failed to create clients: %w", err))
	}

	// Create server
	srvr := server.NewServer(cfg)

	// Register middlewares
	mdlwrs := middleware.NewMiddlewares(cfg)

	// Create handlers
	hdlrs := handlers.NewHandlers(cfg, mdlwrs)
	hdlrs.RegisterRoutes(srvr.Router)

	logger.Info(ctx, cfg.AuthHostPort)

	// Run server
	srvr.Run(ctx)

}
