package main

import (
	"log"
	"project/config"
	"project/internal/api"
	"project/internal/database"
	"project/internal/redis"
)

// @title Messaging API
// @version 1.0
// @description API documentation for the messaging service
// @contact.name Çağla Çolak
// @contact.email caglaccolak@gmail.com
// @host localhost:8080
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	
	// Initialize database connection
	if err := database.InitializeDB(cfg); err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	defer database.CloseDB()

	// Initialize Redis connection
	if err := redis.InitializeRedis(cfg); err != nil {
		log.Fatalf("Error initializing Redis: %v", err)
	}
	defer redis.CloseRedis()

	// Start message scheduler
	
	// Start API server
	if err := api.StartServer(); err != nil {
		log.Fatalf("Error starting API server: %v", err)
	}
}
