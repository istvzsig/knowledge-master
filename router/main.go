package router

import (
	"github.com/gin-gonic/gin"
	"github.com/istvzsig/knowledge-master/handlers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	fm := handlers.NewFAQMaster()

	r.GET("/faqs", fm.HandleFetchFAQs)
	r.POST("/faqs", fm.HandleCreateFAQ)
	r.DELETE("/faqs", fm.HandleDeleteFAQs)

	return r
}
