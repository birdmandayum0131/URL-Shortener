package main

import (
	"errutil"
	"flag"
	"fmt"
	"logger"
	"os"
	"urlshortener/domain"
	"urlshortener/infrastructure"
	"urlshortener/interfaces/repositories"
	"urlshortener/interfaces/rest/routes"
	"urlshortener/services"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v3"

	handlers "urlshortener/interfaces/rest/handlers"
)

var configPath = flag.String("config", "./configs/database.yaml", "config files path")

func main() {
	flag.Parse()

	// * load configs
	dbConfig, err := loadConfigs(*configPath)
	println("Load config success!")
	errutil.PanicIfError(err)

	// * Initialize DB
	dbHandler, err := initDB(dbConfig)
	errutil.PanicIfError(err)
	println("Success to connect to MySQL!")
	defer dbHandler.Conn.Close()

	// * Initialize SnowFlake Node
	node, err := initNode()
	snowflake := domain.SnowFlake{Node: node}
	println("Snowflake init success!")
	errutil.PanicIfError(err)

	// * Initialize restful handler
	handler := createHandler(dbHandler, snowflake)

	// * Setup router
	app := SetupRouter(handler)
	err = app.Run(":8000")
	errutil.PanicIfError(err)
}

func loadConfigs(configPath string) (infrastructure.DBConfig, error) {
	var dbConfig infrastructure.DBConfig

	// * load config files
	cfgFile, err := os.ReadFile(configPath)
	if err != nil {
		return dbConfig, fmt.Errorf("Failed to read database yaml: %v", err)
	}

	// * parse config data to dbConfig
	envCfg := os.ExpandEnv(string(cfgFile))
	err = yaml.Unmarshal([]byte(envCfg), &dbConfig)
	if err != nil {
		return dbConfig, fmt.Errorf("Failed parse config data to db config: %v", err)
	}

	return dbConfig, nil
}

func initDB(config infrastructure.DBConfig) (*infrastructure.MySQLURLDBHandler, error) {
	// * Connect to DB
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := sqlx.Open(config.Driver, dbSource)
	if err != nil {
		return nil, fmt.Errorf("Fail to connect to url db: %v", err)
	}

	// * Create tables
	dbHandler := &infrastructure.MySQLURLDBHandler{
		Conn:   db,
		Logger: &logger.SimpleStdLogger{},
	}
	err = dbHandler.Init()
	if err != nil {
		return nil, fmt.Errorf("Fail to initialize url db: %v", err)
	}
	return dbHandler, nil
}

func initNode() (*snowflake.Node, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, fmt.Errorf("Failed to create snowflake node: %v", err)
	}
	return node, nil
}

func createHandler(dbHandler repositories.URLDBHandler, snowflake domain.SnowFlake) *handlers.URLHandler {
	logger := &logger.SimpleStdLogger{}

	urlRepo := &repositories.URLRepository{DBHandler: dbHandler, Logger: logger}
	hashGen := &domain.SnowFlakeHashGenerator{IDGenerator: snowflake}
	urlItr := services.URLEntryInteractor{URLRepository: urlRepo, HashGenerator: hashGen, Logger: logger}

	return &handlers.URLHandler{URLInteractor: urlItr, Logger: logger}
}

func SetupRouter(handler *handlers.URLHandler) *gin.Engine {
	router := gin.Default()
	routes.InitUrlRoutes(router, handler)
	return router
}
