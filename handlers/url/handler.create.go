package handlers

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type shorten struct {
	ID       int64  `json:"id" db:"id"`
	LongURL  string `json:"longUrl" db:"longURL"`
	ShortURL string `json:"shortUrl" db:"shortURL"`
}

var UrlDB *sqlx.DB
var IdNode *snowflake.Node

// postShorten adds a shorten url from JSON received in the request body.
func CreateUrlHandler(c *gin.Context) {
	var newShorten shorten
	// Call BindJSON to bind the received JSON to
	// newShorten.
	if err := c.Bind(&newShorten); err != nil {
		println(err)
		return
	}

	// * get url id
	results, err := UrlDB.Queryx(fmt.Sprintf("SELECT id, shortURL, longURL FROM url WHERE longURL = '%s'", newShorten.LongURL))
	if err != nil {
		panic(err.Error())
	}

	if results.Next() {
		results.StructScan(&newShorten)
	} else {
		id := IdNode.Generate()
		newShorten.ID = id.Int64()
		var base62Int big.Int
		base62Int.SetInt64(newShorten.ID)
		newShorten.ShortURL = base62Int.Text(62)
		sqlStatement := `
INSERT INTO url (id, shortURL, longURL)
VALUES (:id, :shortURL, :longURL)`
		_, err = UrlDB.NamedExec(sqlStatement, newShorten)
		if err != nil {
			panic(err)
		}
	}

	// Add the new album to the slice.
	c.IndentedJSON(http.StatusCreated, newShorten)
}
