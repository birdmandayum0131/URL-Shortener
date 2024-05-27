package middlewares

import "github.com/gin-gonic/gin"

func URLMiddlewares(r *gin.Engine) []gin.HandlerFunc {
	middlewares := make([]gin.HandlerFunc, 0)
	middlewares = append(middlewares, URLShortenerMiddleware(r))
	return middlewares
}
