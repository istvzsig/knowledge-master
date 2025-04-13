package main

import (
	"github.com/istvzsig/knowledge-master/db"
	"github.com/istvzsig/knowledge-master/pkg/config"
	"github.com/istvzsig/knowledge-master/pkg/router"
)

func main() {
	db.InitFirestore()
	cfg := config.LoadConfig()
	r := router.SetupRouter()

	// Start the server
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		panic("Failed to start server")
	}
}
