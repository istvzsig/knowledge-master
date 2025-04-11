package main

import (
	"log"
	"os"

	"github.com/istvzsig/knowledge-master/db"
	"github.com/istvzsig/knowledge-master/router"
)

func main() {
	db.InitFirestore()
	r := router.InitRouter()

	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
