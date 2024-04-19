package main

import (
	"shortener/routes"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	handlers "shortener/handlers/url"
)

func main() {
	// TODO: remove these init in following refactor
	InitDB()
	InitNode()
	defer handlers.UrlDB.Close()

	app := SetupRouter()
	app.Run("localhost:8000")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	routes.InitUrlRoutes(router)
	return router
}

func InitDB() {
	var err error
	handlers.UrlDB, err = sqlx.Open("mysql", "shortener:shortener@tcp(localhost:3306)/url_shortener")
	if err != nil {
		panic(err.Error())
	}

	println("Success to connect to MySQL!")
}

func InitNode() {
	var err error
	handlers.IdNode, err = snowflake.NewNode(1)
	if err != nil {
		println(err)
		panic(err.Error())
	}
}
