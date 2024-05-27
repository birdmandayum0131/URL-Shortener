package middlewares

import "github.com/gin-gonic/gin"

func URLShortenerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
