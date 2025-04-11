package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/istvzsig/knowledge-master/db"
	"github.com/istvzsig/knowledge-master/models"
)

type FAQMaster struct {
	FAQs map[string]models.FAQ
}

func NewFAQMaster() *FAQMaster {
	return &FAQMaster{
		FAQs: make(map[string]models.FAQ),
	}
}

func (fm *FAQMaster) HandleFetchFAQs(c *gin.Context) {
	faqs, err := fm.fetchFAQsFromDB()
	if err != nil {
		fm.handleError(c, http.StatusInternalServerError, err)
		return
	}

	limit := 100
	paginatedFAQs, err := paginateFAQs(faqs, c.Query("next"), limit)
	if err != nil {
		fm.handleError(c, http.StatusBadRequest, err)
		return
	}

	if len(paginatedFAQs) == 0 {
		fm.handleError(c, http.StatusNotFound, "No more FAQs available")
		return
	}

	c.JSON(http.StatusOK, paginatedFAQs)
}

func (fm *FAQMaster) HandleCreateFAQ(c *gin.Context) {
	var faq models.FAQ
	if err := c.ShouldBindJSON(&faq); err != nil {
		fm.handleError(c, http.StatusBadRequest, err)
		return
	}
	faq.CreatedAt = time.Now().Unix()

	newID, err := fm.createFAQInDB(faq)
	if err != nil {
		fm.handleError(c, http.StatusInternalServerError, err)
		return
	}

	faq.ID = newID
	c.JSON(http.StatusCreated, faq)
}

func (fm *FAQMaster) HandleDeleteFAQs(c *gin.Context) {
	if err := fm.deleteAllFAQsFromDB(); err != nil {
		fm.handleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All FAQs deleted successfully"})
}

func (fm *FAQMaster) fetchFAQsFromDB() ([]models.FAQ, error) {
	ctx := context.Background()
	ref := db.FirestoreClient.NewRef("faqs")

	faqs := fm.FAQs
	if err := ref.Get(ctx, &faqs); err != nil {
		return nil, err
	}

	var faqList []models.FAQ
	for key, faq := range faqs {
		faq.ID = key
		faqList = append(faqList, faq)
	}
	return faqList, nil
}

func (fm *FAQMaster) createFAQInDB(faq models.FAQ) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
	defer cancel()

	ref := db.FirestoreClient.NewRef("faqs")
	newRef, err := ref.Push(ctx, faq)
	if err != nil {
		return "", err
	}
	return newRef.Key, nil
}

func (fm *FAQMaster) deleteAllFAQsFromDB() error {
	ctx := context.Background()
	ref := db.FirestoreClient.NewRef("faqs")
	return ref.Set(ctx, nil)
}

func (fm *FAQMaster) handleError(c *gin.Context, status int, err interface{}) {
	c.JSON(status, gin.H{"error": err})
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
