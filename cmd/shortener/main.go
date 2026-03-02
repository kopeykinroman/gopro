package main

import (
	"fmt"

	"github.com/kopeykinroman/gopro/internal/config"
	"github.com/kopeykinroman/gopro/internal/handlers"
	"github.com/kopeykinroman/gopro/internal/storage"
)

func main() {
	cfg := config.NewConfig()

	store := storage.NewStorage()

	httpHandlers := handlers.NewHTTPHandlers(
		store,
		handlers.WithBaseURL(cfg.BaseURL),
	)
	httpServer := handlers.NewHTTPServer(httpHandlers, cfg.ServerAddress)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Failed to start http server:", err)
	}
}
