package main

import (
	"log"

	"sosmed/config"
	"sosmed/database"
	"sosmed/logger"
	"sosmed/routes"
)

func main() {
	// Load Config
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Logger
	log := logger.NewLogger()

	// Initialize Database
	db, err := database.InitDBMysql()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Initialize Routes
	router := routes.InitRoutes(db, log)

	// Start Server
	if err := router.Run(":8030"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
