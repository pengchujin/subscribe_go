package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello")
		})
	}
	return router
}
