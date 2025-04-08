package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
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
	// Get the current index from the query parameter
	currentIndexStr := c.Query("next") // e.g., /faqs?next=1
	currentIndex := 0
	if currentIndexStr != "" {
		var err error
		currentIndex, err = strconv.Atoi(currentIndexStr)
		if err != nil || currentIndex < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
			return
		}
	}

	// Define the number of FAQs to return per page
	const pageSize = 100

	// Calculate the start and end indices for pagination
	start := currentIndex * pageSize
	end := start + pageSize

	// Check if the start index is within the bounds of the FAQ list
	if start >= len(faqList) {
		c.JSON(http.StatusNotFound, gin.H{"error": "No more FAQs available"})
		return
	}

	// Ensure the end index does not exceed the length of the list
	if end > len(faqList) {
		end = len(faqList)
	}

	// Return the paginated FAQs
	paginatedFAQs := faqList[start:end]
	c.JSON(http.StatusOK, paginatedFAQs)
}

func createFAQ(c *gin.Context) {
	var faq FAQ
	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	faq.CreatedAt = time.Now().Unix()

	ctx := context.Background()
	ref := firestoreClient.NewRef("faqs")

	newRef, err := ref.Push(ctx, faq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	faq.ID = newRef.Key

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

	if err := r.Run(":" + os.Getenv(("PORT"))); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
