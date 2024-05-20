package main

import (
	"fmt"
	"logger"
	"urlshortener/domain"
	"urlshortener/infrastructure"
	"urlshortener/interfaces/repositories"
	"urlshortener/interfaces/rest/routes"
	"urlshortener/services"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	handlers "urlshortener/interfaces/rest/handlers"
)

func main() {
	// TODO: refactor this to config file
	dbConfig := infrastructure.DBConfig{
		Host:     "db",
		Port:     3306,
		User:     "root",
		Password: "password",
		Database: "url_shortener",
		Driver:   "mysql",
	}

	// * Initialize DB
	dbHandler := initDB(dbConfig)
	defer dbHandler.Conn.Close()

	// * Initialize SnowFlake Node
	node := initNode()

	// * Initialize restful handler
	handler := createHandler(dbHandler, node)

	// * Setup router
	app := SetupRouter(handler)

	err := app.Run(":8000")
	if err != nil {
		panic(err.Error())
	}
}

// TODO: maybe we can refactor these setup/init functions with better code style

func SetupRouter(handler *handlers.URLHandler) *gin.Engine {
	router := gin.Default()
	routes.InitUrlRoutes(router, handler)
	return router
}

func initDB(config infrastructure.DBConfig) *infrastructure.MySQLURLDBHandler {
	// * Connect to DB
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := sqlx.Open(config.Driver, dbSource)
	if err != nil {
		panic(err.Error())
	}

	// * Create tables
	dbHandler := &infrastructure.MySQLURLDBHandler{
		Conn:   db,
		Logger: &logger.SimpleStdLogger{},
	}
	dbHandler.Init()

	println("Success to connect to MySQL!")
	return dbHandler
}

func initNode() *snowflake.Node {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err.Error())
	}
	return node
}

func createHandler(dbHandler repositories.URLDBHandler, node *snowflake.Node) *handlers.URLHandler {
	return &handlers.URLHandler{
		URLInteractor: services.URLEntryInteractor{
			URLRepository: &repositories.URLRepository{
				DBHandler: dbHandler,
				Logger:    &logger.SimpleStdLogger{},
			},
			HashGenerator: &domain.SnowFlakeHashGenerator{
				IDGenerator: domain.SnowFlake{Node: node},
			},
			Logger: &logger.SimpleStdLogger{},
		},
		Logger: &logger.SimpleStdLogger{},
	}
}
