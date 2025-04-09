package config

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var FirestoreClient *db.Client

func InitFirestore() {
	ctx := context.Background()
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	config := &firebase.Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	FirestoreClient, err = app.Database(ctx)
	if err != nil {
		log.Fatalf("Error getting Database client: %v", err)
	}
}
