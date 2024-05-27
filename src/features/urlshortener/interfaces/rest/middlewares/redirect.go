package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func URLShortenerMiddleware(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		hash := c.Param("hash")
		c.Request.URL.Path = fmt.Sprintf("/api/v1/url/%s", hash)
		r.HandleContext(c)
		c.Next()
	}
}
