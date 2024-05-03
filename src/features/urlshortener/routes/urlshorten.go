package routes

import (
	handlers "urlshortener/handlers"

	"github.com/gin-gonic/gin"
)

// Initialize route for CRUD operation of url shorten
func InitUrlRoutes(route *gin.Engine, handler *handlers.URLHandler) {
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/url/shorten", handler.CreateURLHandler)
	groupRoute.GET("/url/:hash", handler.GetURLHandler)
}
