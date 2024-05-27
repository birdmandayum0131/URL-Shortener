package routes

import (
	handlers "urlshortener/interfaces/rest/handlers"

	"github.com/gin-gonic/gin"
)

// Initialize route for CRUD operation of url shorten
func InitUrlRoutes(route *gin.Engine, middlewares []gin.HandlerFunc, handler *handlers.URLHandler) {
	groupRoute := route.Group("/api/v1")
	groupRoute.Use(middlewares...)
	groupRoute.POST("/url/shorten", handler.CreateURLHandler)
	groupRoute.GET("/url/:hash", handler.GetURLHandler)
}
