package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"task-management-system/config"
	"task-management-system/internal/server"
	"time"

	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()

	cfg.Logger.Info("Logger initialized")

	s := server.NewServer(cfg)

	go func() {
		if err := s.Start(); err != nil {
			log.Fatalf("Server failed to start: %v\n", err)
			cfg.Logger.Error("Server failed to start", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Stop(); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}

	log.Println("Server shutdown gracefully")
}
