package middleware

import (
	"github.com/gin-gonic/gin"
	"test10/pkg/utils"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 400
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 // token无权限
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //token无效或过期

			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
