package main

import (
	"context"
	"github.com/anger-aa/quotes/internal/handler"
	"github.com/anger-aa/quotes/internal/storage"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	storage := storage.NewStorage()
	handler := handler.NewHandler(storage)
	srv := NewServer()

	go func() {
		if err := srv.Run(":8080", handler.InitRoutes()); err != nil {
			log.Fatalf("Error running the server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error shutting down the server: %s", err)
	}
}
