package main

import (
	"log"
	"os"

	"github.com/istvzsig/knowledge-master/db"
	"github.com/istvzsig/knowledge-master/handlers"
	"github.com/istvzsig/knowledge-master/router"
)

func main() {
	db.InitFirestore()

	port := os.Getenv("BACKEND_PORT")
	r := router.NewRouter(":" + port)

	fm := handlers.NewFAQMaster()
	r.GET("/faqs", fm.HandleFetchFAQs)
	r.POST("/faqs", fm.HandleCreateFAQ)
	r.DELETE("/faqs", fm.HandleDeleteFAQs)

	if err := r.Run(r.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
