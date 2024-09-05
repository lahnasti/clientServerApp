package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/lahnasti/clientServerApp/server/config"
	"github.com/lahnasti/clientServerApp/server/handlers"
	"github.com/lahnasti/clientServerApp/server/logger"
	"github.com/lahnasti/clientServerApp/server/repository"
	"github.com/lahnasti/clientServerApp/server/routes"
)

func main() {

	fmt.Println("Server starting")
	cfg := config.ReadConfig()
	fmt.Println(cfg)
	zlog := logger.SetupLogger(cfg.DebugFlag)
	zlog.Debug().Any("config", cfg).Msg("Check cfg value")
	zlog.Debug().Str("migration_path", cfg.MPath).Msg("Path to migrations")

	conn, err := initDB(cfg.DBAddr)
	if err != nil {
		zlog.Fatal().Err(err).Msg("Connection DB failed")
	}

	err = repository.Migrations(cfg.DBAddr, cfg.MPath, zlog)
	if err != nil {
		zlog.Fatal().Err(err).Msg("Init migrations failed")
	}

	db, err := repository.NewDB(conn)
	if err != nil {
		zlog.Fatal().Err(err).Msg("Unable to create database storage")
	}

	server := handlers.NewServer(db)

	r := routes.SetupRoutes(server)

	zlog.Info().Msgf("Starting server on %s", cfg.Addr)

	if err := r.Run(cfg.Addr); err != nil {
		zlog.Fatal().Err(err).Msg("Failed to start server")
	}
}

func initDB(addr string) (*pgx.Conn, error) {
	for i := 0; i < 7; i++ {
		time.Sleep(2 * time.Second)
		conn, err := pgx.Connect(context.Background(), addr)
		if err == nil {
			return conn, nil
		}
	}
	conn, err := pgx.Connect(context.Background(), addr)
	if err != nil {
		return nil, fmt.Errorf("database initialization error: %w", err)
	}
	return conn, nil
}
