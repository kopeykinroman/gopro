package main

import (
	"fmt"

	"github.com/kopeykinroman/gopro/internal/handlers"
	"github.com/kopeykinroman/gopro/internal/storage"
)

func main() {
	store := storage.NewStorage()
	httpHandlers := handlers.NewHTTPHandlers(store)
	httpServer := handlers.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Failed to start http server:", err)
	}
}
