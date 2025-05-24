package main

import (
	"log"

	"sosmed/config"
	"sosmed/database"
	"sosmed/logger"
	"sosmed/src/interactions"
	"sosmed/src/notifications"
	posts "sosmed/src/posts"
	users "sosmed/src/users"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	
	client := &websocket.Conn{}

	// Initialize Gin Router
	router := gin.Default()

	// Register User Routes
	users.RegisterRoutes(router, db, log)
	posts.RegisterRoutes(router, db, log)
	interactions.RegisterRoutes(router, db, log)
	notifications.RegisterRoutes(router, db, log, client)

	// Start Server
	port := ":8888"
	log.Infof("Starting server on %s", port)
	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
