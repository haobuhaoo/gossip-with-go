package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "production" {
		_ = godotenv.Load()
	}

	ctx := context.Background()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	cfg := config{
		addr: ":" + port,
		db: dbConfig{
			dsn: os.Getenv("GOOSE_DBSTRING"),
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	poolConfig, err := pgxpool.ParseConfig(cfg.db.dsn)
	if err != nil {
		slog.Error("Invalid DB DSN", "error", err)
		os.Exit(1)
	}

	poolConfig.MaxConns = 25
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = 30 * time.Minute
	poolConfig.MaxConnIdleTime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	slog.Info("Connected to database")

	api := application{
		config: cfg,
		db:     pool,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
