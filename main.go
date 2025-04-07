package main

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"google.golang.org/api/option"
)

type FAQ struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
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

func getFAQs(c *gin.Context) {
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

func addFAQ(c *gin.Context) {
	var faq FAQ
	faq.ID = "asdasd"
	faq.Question = "Question?"
	faq.Answer = "Anwers!"
	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	ref := firestoreClient.NewRef("faqs")

	// Push the new FAQ and handle both return values
	newRef, err := ref.Push(ctx, faq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set the ID of the FAQ to the newly created document's key
	faq.ID = newRef.Key

	c.JSON(http.StatusCreated, faq)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	initFirestore()

	r := gin.Default()
	r.GET("/faqs", getFAQs)
	r.POST("/faqs", addFAQ)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
