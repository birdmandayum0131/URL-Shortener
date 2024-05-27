package main

import (
	"errors"
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
	"gopkg.in/yaml.v3"

	handlers "urlshortener/interfaces/rest/handlers"
)

var configPath = flag.String("config", "./configs/database.yaml", "config files path")

func main() {
	flag.Parse()

	// * load configs
	dbConfig, poolConfig, err := loadConfigs(*configPath)
	println("Load config success!")
	errutil.PanicIfError(err)

	// * Initialize DB
	dbHandler, err := initDB(dbConfig, poolConfig)
	errutil.PanicIfError(err)
	println("Success to connect to MySQL!")
	defer dbHandler.Close()

	// * Initialize SnowFlake Node
	node, err := initNode()
	snowflake := domain.SnowFlake{Node: node}
	println("Snowflake init success!")
	errutil.PanicIfError(err)

	// * Initialize restful handler
	handler := createHandler(dbHandler, snowflake)

	// * Setup router
	app := SetupRouter(nil, handler)
	err = app.Run(":8000")
	errutil.PanicIfError(err)
}

func loadConfigs(configPath string) (infrastructure.DBConfig, infrastructure.PoolConfig, error) {
	var dbConfig infrastructure.DBConfig
	var poolConfig infrastructure.PoolConfig

	// * load config files
	cfgFile, err := os.ReadFile(configPath)
	if err != nil {
		return dbConfig, poolConfig, fmt.Errorf("Failed to read database yaml: %v", err)
	}

	// * parse config data to dbConfig
	envCfg := os.ExpandEnv(string(cfgFile))
	errConn := yaml.Unmarshal([]byte(envCfg), &dbConfig)
	errPool := yaml.Unmarshal([]byte(envCfg), &poolConfig)
	err = errors.Join(errConn, errPool)
	if err != nil {
		return dbConfig, poolConfig, fmt.Errorf("Failed parse config data to db config: %v", err)
	}

	return dbConfig, poolConfig, nil
}

func initDB(dbConfig infrastructure.DBConfig, poolConfig infrastructure.PoolConfig) (*infrastructure.MySQLURLDBHandler, error) {
	dbHandler := &infrastructure.MySQLURLDBHandler{Logger: &logger.SimpleStdLogger{}}
	err := dbHandler.Init(dbConfig, poolConfig)
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

func SetupRouter(middlewares []gin.HandlerFunc, handler *handlers.URLHandler) *gin.Engine {
	router := gin.Default()
	routes.InitUrlRoutes(router, middlewares, handler)
	return router
}
