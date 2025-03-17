package main

import (
	"log"
	"os"
	"os/signal"
	"sports-go/cmd/user/config"
	"sports-go/cmd/user/proto"
	"sports-go/cmd/user/server"
	"syscall"
)

func main() {
	cfg := config.LoadConfig()
	user := &proto.User{}
	server := server.NewServer(cfg, user)
	go func() {
		log.Printf("Starting gRPC server on %s...", cfg.GRPCPort)
		if err := server.Start(); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down gRPC server...")
	server.Stop()
}
