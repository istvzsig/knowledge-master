package router

import (
	"github.com/gin-gonic/gin"
	"github.com/istvzsig/knowledge-master/internal/faq"
)

func SetupRouter(service faq.Service) *gin.Engine {
	router := gin.Default()

	api := router.Group("/faqs")
	{
		api.GET("", func(c *gin.Context) { faq.HandleFetchFAQs(c, service) })
		api.POST("", func(c *gin.Context) { faq.HandleCreateFAQ(c, service) })
		api.PATCH("/:id", func(c *gin.Context) { faq.HandleDeleteFAQByID(c, service) })
		api.DELETE("", func(c *gin.Context) { faq.HandleDeleteAllFAQs(c, service) })
	}

	return router
}
