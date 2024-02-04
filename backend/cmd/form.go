package main

import (
	"backend/internal/app/server"
	"backend/internal/config"
	"context"
	"log"
)

func main() {
	logger := log.Default()
	cfg := config.Prod()
	srv := server.NewServer(context.Background(), cfg, logger)
	srv.Start()
}
