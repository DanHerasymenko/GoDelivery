package postgresClient

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DanHerasymenko/GoDelivery/services/auth-service/internal/config"
	"github.com/DanHerasymenko/GoDelivery/shared/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type Client struct {
	Postgres *pgxpool.Pool
}

func NewPostgresClient(ctx context.Context, cfg *config.Config) (*Client, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.AuthPostgresUser,
		cfg.AuthPostgresPassword,
		cfg.PostgresContainerHost,
		cfg.PostgresContainerPort,
		cfg.AuthPostgresDB,
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	// check connection
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping pool: %w", err)
	}

	logger.Info(ctx, "Postgres ping successful")

	if cfg.RunMigrations {
		if err := runMigrations(ctx, dsn); err != nil {
			return nil, err
		}
		logger.Info(ctx, "Postgres migrations successful")
	} else {
		logger.Info(ctx, "Postgres migrations skipped")
	}

	return &Client{Postgres: pool}, nil
}

func runMigrations(ctx context.Context, dsn string) error {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database for migrations: %w", err)
	}
	defer func() {
		if cl := db.Close(); cl != nil {
			logger.Info(ctx, fmt.Sprintf("failed to close migration db connection: %v", cl))
		}
	}()

	if err := goose.Up(db, "./migrations"); err != nil {
		return fmt.Errorf("failed to run goose migrations: %w", err)
	}

	return nil
}
