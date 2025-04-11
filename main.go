package main

import (
	"log"
	"os"

	"github.com/istvzsig/knowledge-master/config"
	"github.com/istvzsig/knowledge-master/routes"
)

func main() {
	config.InitFirestore()

	r := routes.SetupRouter()

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
