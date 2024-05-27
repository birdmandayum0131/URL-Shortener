package middlewares

import "github.com/gin-gonic/gin"

func URLMiddlewares() []gin.HandlerFunc {
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, URLShortenerMiddleware())
	return middlewares
}
