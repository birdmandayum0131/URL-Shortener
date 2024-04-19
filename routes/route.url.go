package routes

import (
	handlers "shortener/handlers/url"

	"github.com/gin-gonic/gin"
)

// Initialize route for CRUD operation of url shorten
func InitUrlRoutes(route *gin.Engine) {
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/url/shorten", handlers.CreateUrlHandler)
	groupRoute.GET("/url/:hash", handlers.GetUrlHandler)
}
