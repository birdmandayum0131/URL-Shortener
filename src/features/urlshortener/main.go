package main

import (
	"logger"
	"urlshortener/domain"
	"urlshortener/infrastructure"
	"urlshortener/interfaces/repositories"
	"urlshortener/interfaces/rest/routes"
	"urlshortener/services"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	handlers "urlshortener/interfaces/rest/handlers"
)

func main() {
	// TODO: remove these init in following refactor
	db := initDB()
	defer db.Close()

	node := initNode()
	handler := initHandler(db, node)

	app := SetupRouter(handler)
	app.Run("localhost:8000")
}

func SetupRouter(handler *handlers.URLHandler) *gin.Engine {
	router := gin.Default()
	routes.InitUrlRoutes(router, handler)
	return router
}

func initDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:password@tcp(localhost:3306)/url_shortener")
	if err != nil {
		panic(err.Error())
	}

	var schema = `
CREATE TABLE url_shortener.urlmappings (
	id BIGINT NOT NULL AUTO_INCREMENT,
	shortURL VARCHAR(64) NULL,
	longURL VARCHAR(2048) NULL,
	PRIMARY KEY (id))`

	_, err = db.Exec(schema)
	if err != nil && !(&mysql.MySQLError{Number: 1050}).Is(err) {
		panic(err.Error())
	}

	println("Success to connect to MySQL!")
	return db
}

func initNode() *snowflake.Node {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err.Error())
	}
	return node
}

func initHandler(db *sqlx.DB, node *snowflake.Node) *handlers.URLHandler {
	return &handlers.URLHandler{
		URLInteractor: services.URLEntryInteractor{
			URLRepository: &repositories.URLRepository{
				DBHandler: &infrastructure.MySQLURLDBHandler{
					Conn:   db,
					Logger: &logger.SimpleStdLogger{},
				},
				Logger: &logger.SimpleStdLogger{},
			},
			HashGenerator: &domain.SnowFlakeHashGenerator{
				IDGenerator: domain.SnowFlake{Node: node},
			},
			Logger: &logger.SimpleStdLogger{},
		},
		Logger: &logger.SimpleStdLogger{},
	}
}
