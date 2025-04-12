package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// HandleError sends a JSON response with the error message and status code
func HandleError(c *gin.Context, status int, err any) {
	log.Printf("Error: %v", err)
	c.JSON(status, gin.H{"error": err})
}
