package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"google.golang.org/api/option"
)

type FAQ struct {
	ID        string `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	CreatedAt int64  `json:"createdAt"`
}

var firestoreClient *db.Client

func initFirestore() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./json/knowledge-manager.json")
	// Load the Firebase configuration
	config := &firebase.Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	firestoreClient, err = app.Database(ctx)
	if err != nil {
		log.Fatalf("error getting Database client: %v\n", err)
	}
}

func fetchFAQs(c *gin.Context) {
	ctx := context.Background()
	ref := firestoreClient.NewRef("faqs")
	var faqs map[string]FAQ

	// Get all FAQs
	if err := ref.Get(ctx, &faqs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert map to slice and add IDs
	var faqList []FAQ
	for key, faq := range faqs {
		faq.ID = key
		faqList = append(faqList, faq)
	}

	c.JSON(http.StatusOK, faqList)
}

func createFAQ(c *gin.Context) {
	var faq FAQ
	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	ref := firestoreClient.NewRef("faqs")

	newRef, err := ref.Push(ctx, faq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	faq.ID = newRef.Key
	faq.CreatedAt = time.Now().Unix()

	c.JSON(http.StatusCreated, faq)
}

func deleteAllFAQs(c *gin.Context) {
	ctx := context.Background()
	ref := firestoreClient.NewRef("faqs")

	// Set the "faqs" node to nil to delete all FAQs
	if err := ref.Set(ctx, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All FAQs deleted successfully"})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	initFirestore()

	r := gin.Default()
	r.GET("/faqs", fetchFAQs)
	r.POST("/faqs", createFAQ)
	r.DELETE("/faqs", deleteAllFAQs)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
