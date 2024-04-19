package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUrlHandler(c *gin.Context) {
	var shortenURL shorten
	hash := c.Param("hash")

	// * get url by hash
	results, err := UrlDB.Queryx(fmt.Sprintf("SELECT id, shortURL, longURL FROM url WHERE shortURL = '%s'", hash))
	if err != nil {
		panic(err.Error())
	}

	if results.Next() {
		results.StructScan(&shortenURL)
		c.IndentedJSON(http.StatusFound, shortenURL)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "shorten not found"})
	}

}