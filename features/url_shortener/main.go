package main

import (
	"logger"
	"shortener/domain"
	"shortener/infrastructure"
	"shortener/repositories"
	"shortener/routes"
	"shortener/services"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	handlers "shortener/handlers"
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
	db, err := sqlx.Open("mysql", "shortener:shortener@tcp(localhost:3306)/url_shortener")
	if err != nil {
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
