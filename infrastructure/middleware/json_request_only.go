package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireJsonRequestBodyOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		if c.Request.Header.Get("Content-Type") != "application/json" {
			c.AbortWithStatus(400)
			return
		}
		c.Next()
	}
}
