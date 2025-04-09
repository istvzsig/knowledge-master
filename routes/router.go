package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/istvzsig/knowledge-menager/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	faqManager := handlers.NewFAQManager()

	r.GET("/faqs", faqManager.HandleFetchFAQs)
	r.POST("/faqs", faqManager.HandleCreateFAQ)
	r.DELETE("/faqs", faqManager.HandleDeleteFAQs)

	return r
}
