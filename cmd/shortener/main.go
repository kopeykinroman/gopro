package main

import (
	"fmt"
	handler "shortener/internal/handler"
	"shortener/internal/storage"
)

func main() {
	store := storage.NewStorage()
	httpHandlers := handler.NewHTTPHandlers(store)
	httpServer := handler.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Failed to start http server:", err)
	}
}
