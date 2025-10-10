package server

import (
	"context"
	"homeworkMtsGo/homework2/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)



func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /version", handlers.VersionHandler)
	mux.HandleFunc("POST /decode", handlers.DecodeHandler)
	mux.HandleFunc("GET /hard-op", handlers.HardOpHandler)

	server := &http.Server{Addr: ":8081", Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

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