package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/pengchujin/subscribe_go/controllers"
	"github.com/pengchujin/subscribe_go/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		user := new(controllers.User)
		v1.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello")
		})
		v1.POST("/signup", user.Store)
		v1.POST("/signin", user.Get)
	}
	test := router.Group("/api/test")
	test.Use(jwt.JWT())
	{
		test.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello")
		})
	}
	return router
}
