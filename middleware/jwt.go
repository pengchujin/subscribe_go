package jwt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pengchujin/subscribe_go/models"
	"github.com/pengchujin/subscribe_go/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		code := "SUCCESS"
		result := models.Result{
			Code:    http.StatusUnauthorized,
			Message: "无法认证，重新登录",
			Data:    nil,
		}
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"result": result,
			})
		}
		token := strings.Fields(auth)[1]
		if token == "" {
			code = "请求参数错误"
		} else {
			claims, err := util.ParseToken(token)
			// Todo
			fmt.Println(claims)
			if err != nil {
				code = "Token鉴权失败"
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = "Token已超时"
			}
		}

		if code != "SUCCESS" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "验证失败",
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
