package knowledge_master

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/istvzsig/knowledge-master/internal/types"
	respTypes "github.com/istvzsig/knowledge-master/pkg/types"
)

// HandleFetchFAQs handles the request to fetch all FAQs.
func HandleFetchFAQs(c *gin.Context) {
	faqs, err := FetchFAQs()
	if err != nil {
		errResp := respTypes.HttpResponseError{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}

		HandleError(c, errResp)
		return
	}
	response := respTypes.HttpResponseSuccess{
		Status:  http.StatusOK,
		Message: "Successfully fetched FAQs.",
		Data:    faqs,
	}
	HandleSuccess(c, response)
}

// HandleCreateFAQ handles the request to create a new FAQ.
func HandleCreateFAQ(c *gin.Context) {
	var newFAQ = &types.FAQ{}

	if err := c.ShouldBindJSON(&newFAQ); err != nil {
		errResp := respTypes.HttpResponseError{
			Status:  http.StatusBadRequest,
			Message: "Invalid FAQ data",
		}
		HandleError(c, errResp)
		return
	}

	faqID, err := CreateFAQ(*newFAQ)
	if err != nil {
		errResp := respTypes.HttpResponseError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Error creating FAQ: %v", err),
		}
		HandleError(c, errResp)
		return
	}

	response := respTypes.HttpResponseSuccess{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("FAQ successfully created with id: %v", faqID),
		Data:    faqID,
	}
	HandleSuccess(c, response)
}

// HandleDeleteAllFAQs handles the request to delete all FAQs.
func HandleDeleteAllFAQs(c *gin.Context) {
	err, _ := DeleteAllFAQs()
	if err != nil {
		errResp := respTypes.HttpResponseError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Error deleting all FAQs: %v", err),
		}
		HandleError(c, errResp)
		return
	}

	response := respTypes.HttpResponseSuccess{
		Status:  http.StatusOK,
		Message: "All FAQs deleted successfully",
	}
	HandleSuccess(c, response)
}

// HandleDeleteFAQByID handles the request to delete a single FAQ by ID.
func HandleDeleteFAQByID(c *gin.Context) {
	faqID := c.Param("id")
	if faqID == "" {
		errResp := respTypes.HttpResponseError{
			Status:  http.StatusInternalServerError,
			Message: "FAQ ID is required",
		}
		HandleError(c, errResp)
		return
	}

	err := DeleteFAQByID(faqID)
	if err != nil {
		errResp := respTypes.HttpResponseError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Error deleting FAQ with ID %s: %v", faqID, err),
		}
		HandleError(c, errResp)
		return
	}
	response := respTypes.HttpResponseSuccess{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("FAQ with id: %v successfully deleted.", faqID),
	}
	HandleSuccess(c, response)
}

// HandleError sends a JSON response with the error message and status code
func HandleError(c *gin.Context, resp respTypes.HttpResponseError) {
	c.JSON(resp.Status, resp)
}

// HandleSuccess sends a JSON response with the status code and data
func HandleSuccess(c *gin.Context, resp respTypes.HttpResponseSuccess) {
	c.JSON(resp.Status, resp)
}
