package main

import (
	"context"
	"fmt"
	"homework3/internal/application/usecase"
	"homework3/internal/infrastructure/api"
	"homework3/internal/infrastructure/config"
	"homework3/internal/infrastructure/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.Load()
	if (err != nil) {
		log.Fatal(err)
	}
	repository, err := db.GetRepository(cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName)
	if (err != nil) {
		log.Fatal(err)
	}
	service := usecase.NewUserService(repository)
	handlers := api.NewHandler(service)
	addr := fmt.Sprintf(":%s", cfg.HTTPPort)
	server := &http.Server{
		Addr:    addr,
		Handler: handlers.SetupRoutes(),
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	<-osSignals
	log.Println("Graceful shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited properly")
}