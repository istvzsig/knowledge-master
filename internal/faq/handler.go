package faq

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/istvzsig/knowledge-master/internal/types"
	"github.com/istvzsig/knowledge-master/pkg/utils"
)

// HandleFetchFAQs handles the request to fetch all FAQs.
func HandleFetchFAQs(c *gin.Context, service Service) {
	faqs, err := service.FetchFAQs()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, fmt.Sprintf("Error fetching FAQs: %v", err))
		return
	}
	c.JSON(http.StatusOK, faqs)
}

// HandleCreateFAQ handles the request to create a new FAQ.
func HandleCreateFAQ(c *gin.Context, service Service) {
	var newFAQ = &types.FAQ{}

	if err := c.ShouldBindJSON(&newFAQ); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid FAQ data")
		return
	}

	faqID, err := service.CreateFAQ(*newFAQ)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, fmt.Sprintf("Error creating FAQ: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "FAQ created successfully", "id": faqID})
}

// HandleDeleteAllFAQs handles the request to delete all FAQs.
func HandleDeleteAllFAQs(c *gin.Context, service Service) {
	err := service.DeleteAllFAQs()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, fmt.Sprintf("Error deleting all FAQs: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All FAQs deleted successfully"})
}

// HandleDeleteFAQByID handles the request to delete a single FAQ by ID.
func HandleDeleteFAQByID(c *gin.Context, service Service) {
	faqID := c.Param("id")
	if faqID == "" {
		utils.HandleError(c, http.StatusNotFound, "FAQ ID is required")
		return
	}

	err := service.DeleteFAQByID(faqID)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, fmt.Sprintf("Error deleting FAQ with ID %s: %v", faqID, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("FAQ with ID %s deleted successfully", faqID)})
}
