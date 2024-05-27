package routes

import (
	handlers "urlshortener/interfaces/rest/handlers"
	"urlshortener/interfaces/rest/middlewares"

	"github.com/gin-gonic/gin"
)

// Initialize route for CRUD operation of url shorten
func InitUrlRoutes(route *gin.Engine, urlMiddlewares []gin.HandlerFunc, handler *handlers.URLHandler) {
	route.GET("/:hash", middlewares.URLShortenerMiddleware(route))
	groupRoute := route.Group("/api/v1")
	groupRoute.Use(urlMiddlewares...)
	groupRoute.POST("/url/shorten", handler.CreateURLHandler)
	groupRoute.GET("/url/:hash", handler.GetURLHandler)
}
