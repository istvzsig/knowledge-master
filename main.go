package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/istvzsig/knowledge-menager/config"
	"github.com/istvzsig/knowledge-menager/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.InitFirestore()

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
