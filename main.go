package main

import (
	"github.com/istvzsig/knowledge-master/db"
	"github.com/istvzsig/knowledge-master/internal/faq"
	"github.com/istvzsig/knowledge-master/internal/types"
	"github.com/istvzsig/knowledge-master/pkg/config"
	"github.com/istvzsig/knowledge-master/pkg/router"
)

func main() {
	db.InitFirestore()
	cfg := config.LoadConfig()
	km := types.NewKnowledgeMaster()
	s := faq.NewFAQService(km)
	r := router.SetupRouter(s)

	// Start the server
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		panic("Failed to start server")
	}
}
