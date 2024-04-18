package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/big"
	"net/http"
)

type shorten struct {
	ID       int64  `json:"id" db:"id"`
	LongURL  string `json:"longUrl" db:"longURL"`
	ShortURL string `json:"shortUrl" db:"shortURL"`
}

var db *sqlx.DB
var node *snowflake.Node

// // getAlbums responds with the list of all albums as JSON.
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// postShorten adds a shorten url from JSON received in the request body.
func postShorten(c *gin.Context) {
	var newShorten shorten
	// Call BindJSON to bind the received JSON to
	// newShorten.
	if err := c.Bind(&newShorten); err != nil {
		println(err)
		return
	}

	// * get url id
	results, err := db.Queryx(fmt.Sprintf("SELECT id, shortURL, longURL FROM url WHERE longURL = '%s'", newShorten.LongURL))
	if err != nil {
		panic(err.Error())
	}

	if results.Next() {
		results.StructScan(&newShorten)
	} else {
		id := node.Generate()
		newShorten.ID = id.Int64()
		var base62Int big.Int
		base62Int.SetInt64(newShorten.ID)
		newShorten.ShortURL = base62Int.Text(62)
		sqlStatement := `
INSERT INTO url (id, shortURL, longURL)
VALUES (:id, :shortURL, :longURL)`
		_, err = db.NamedExec(sqlStatement, newShorten)
		if err != nil {
			panic(err)
		}
	}

	// Add the new album to the slice.
	c.IndentedJSON(http.StatusCreated, newShorten)
}

// // getAlbumByID locates the album whose ID value matches the id
// // parameter sent by the client, then returns that album as a response.
func getUrlByHash(c *gin.Context) {
	var shortenURL shorten
	hash := c.Param("hash")

	// * get url by hash
	results, err := db.Queryx(fmt.Sprintf("SELECT id, shortURL, longURL FROM url WHERE shortURL = '%s'", hash))
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

func main() {
	var err error

	db, err = sqlx.Open("mysql", "shortener:shortener@tcp(localhost:3306)/url_shortener")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	println("Success to connect to MySQL!")

	node, err = snowflake.NewNode(1)
	if err != nil {
		println(err)
		panic(err.Error())
	}

	router := gin.Default()
	router.POST("/shorten", postShorten)
	router.GET("/shortUrl/:hash", getUrlByHash)

	router.Run("localhost:8000")
}
