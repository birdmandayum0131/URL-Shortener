package handlers

import (
	"fmt"
	"net/http"
	schemas "urlshortener/interfaces/schemas/api"
	"urlshortener/services"

	"github.com/gin-gonic/gin"
)

// class to handle url shorten tasks
type URLHandler struct {
	URLInteractor services.URLEntryInteractor
}

// Create a new url entry
func (handler *URLHandler) CreateURLHandler(c *gin.Context) {
	var request schemas.CreateURLRequest
	// Call BindJSON to bind the received JSON request
	if err := c.Bind(&request); err != nil {
		msg := fmt.Sprintf("Invalid request format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	shortURL, err := handler.URLInteractor.CreateEntry(request.LongURL)
	if err != nil {
		msg := fmt.Sprintf("Internal server error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	response := schemas.CreateURLResponse{
		LongURL:  request.LongURL,
		ShortURL: shortURL,
	}

	c.IndentedJSON(http.StatusCreated, response)
}

func (handler *URLHandler) GetURLHandler(c *gin.Context) {
	shortURL := c.Param("hash")

	longURL, err := handler.URLInteractor.GetURL(shortURL)
	if err != nil {
		msg := fmt.Sprintf("Internal server error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	c.Redirect(http.StatusFound, longURL)
}
