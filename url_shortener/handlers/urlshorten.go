package handlers

import (
	"logger"
	"net/http"
	schemas "shortener/schemas/api"
	"shortener/services"

	"github.com/gin-gonic/gin"
)

// class to handle url shorten tasks
type URLHandler struct {
	URLInteractor services.URLEntryInteractor
	Logger        logger.Logger
}

// Create a new url entry
func (handler *URLHandler) CreateURLHandler(c *gin.Context) {
	var request schemas.CreateURLRequest
	// Call BindJSON to bind the received JSON request
	if err := c.Bind(&request); err != nil {
		handler.Logger.Log(err.Error())
		return
	}

	shortURL, err := handler.URLInteractor.CreateEntry(request.LongURL)
	if err != nil {
		handler.Logger.Log(err.Error())
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
		handler.Logger.Log(err.Error())
		return
	}

	response := schemas.GetURLResponse{
		LongURL:  longURL,
		ShortURL: shortURL,
	}

	c.IndentedJSON(http.StatusFound, response)
}
