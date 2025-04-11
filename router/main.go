package router

import (
	"github.com/gin-gonic/gin"
)

type router struct {
	*gin.Engine
	Port string
}

func NewRouter(port string) *router {
	return &router{
		Engine: gin.Default(),
		Port:   port,
	}
}
