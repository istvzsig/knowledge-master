package router

import (
	"github.com/gin-gonic/gin"
	"github.com/istvzsig/knowledge-master/internal/knowledge_master"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/faqs")
	{
		api.GET("", func(c *gin.Context) { knowledge_master.HandleFetchFAQs(c) })
		api.POST("", func(c *gin.Context) { knowledge_master.HandleCreateFAQ(c) })
		api.PATCH("/:id", func(c *gin.Context) { knowledge_master.HandleDeleteFAQByID(c) })
		api.DELETE("", func(c *gin.Context) { knowledge_master.HandleDeleteAllFAQs(c) })
	}

	return r
}
