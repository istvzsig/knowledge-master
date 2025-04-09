package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/istvzsig/knowledge-menager/config"
	"github.com/istvzsig/knowledge-menager/models"

	"github.com/gin-gonic/gin"
)

type FAQManager struct {
	FAQs map[string]models.FAQ
}

func NewFAQManager() *FAQManager {
	return &FAQManager{
		FAQs: make(map[string]models.FAQ),
	}
}

func (fm *FAQManager) HandleFetchFAQs(c *gin.Context) {
	ctx := context.Background()
	ref := config.FirestoreClient.NewRef("faqs")

	faqs := fm.FAQs
	if err := ref.Get(ctx, &faqs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var faqList []models.FAQ
	for key, faq := range faqs {
		faq.ID = key
		faqList = append(faqList, faq)
	}

	paginatedFAQs, err := paginateFAQs(faqList, c.Query("next"), 100)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(paginatedFAQs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No more FAQs available"})
		return
	}

	c.JSON(http.StatusOK, paginatedFAQs)
}

func (fm *FAQManager) HandleCreateFAQ(c *gin.Context) {
	var faq models.FAQ
	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	faq.CreatedAt = time.Now().Unix()

	ctx := context.Background()
	ref := config.FirestoreClient.NewRef("faqs")

	newRef, err := ref.Push(ctx, faq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	faq.ID = newRef.Key
	c.JSON(http.StatusCreated, faq)
}

func (fm *FAQManager) HandleDeleteFAQs(c *gin.Context) {
	ctx := context.Background()
	ref := config.FirestoreClient.NewRef("faqs")

	if err := ref.Set(ctx, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All FAQs deleted successfully"})
}

func paginateFAQs(faqs []models.FAQ, indexStr string, pageSize int) ([]models.FAQ, error) {
	startIndex := 0
	if indexStr != "" {
		i, err := strconv.Atoi(indexStr)
		if err != nil || i < 0 {
			return nil, err
		}
		startIndex = i * pageSize
	}

	endIndex := startIndex + pageSize
	if startIndex >= len(faqs) {
		return []models.FAQ{}, nil
	}
	if endIndex > len(faqs) {
		endIndex = len(faqs)
	}
	return faqs[startIndex:endIndex], nil
}
