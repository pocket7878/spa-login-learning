package middleware

import "github.com/gin-gonic/gin"

func RequireJsonRequestBodyOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			c.AbortWithStatus(400)
			return
		}
		c.Next()
	}
}
