package controllers

import (
	"github.com/gin-gonic/gin"
)

type Basic struct {
	
}

func (basic *Basic) JsonSuccess (c *gin.Context, status int, h gin.H) {
	h["status"] = "success"
	c.JSON(status, h)
	return
}

