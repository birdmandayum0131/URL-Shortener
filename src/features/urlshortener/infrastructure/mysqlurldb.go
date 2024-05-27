package infrastructure

import (
	"dbutil"
	"fmt"
	"logger"
	"time"
	"urlshortener/interfaces/models"
	"urlshortener/interfaces/schemas"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLURLDBHandler struct {
	conn   *sqlx.DB
	Logger logger.Logger
}

// Initialize mysql url database
func (dbHandler *MySQLURLDBHandler) Init(dbConfig DBConfig, poolConfig PoolConfig) error {
	// * load configs
	err := dbHandler.loadDBConfig(dbConfig)
	if err != nil {
		return fmt.Errorf("Fail to init db: %v", err)
	}
	dbHandler.loadPoolConfig(poolConfig)

	// * create url table
	// TODO: refactor the sql statement to better write style
	var schema = `
	CREATE TABLE url_shortener.urlmappings (
		id BIGINT NOT NULL AUTO_INCREMENT,
		shortURL VARCHAR(64) NULL,
		longURL VARCHAR(2048) NULL,
		PRIMARY KEY (id))`

	_, err = dbHandler.conn.Exec(schema)
	// * check if table is created
	if err != nil && !(&mysql.MySQLError{Number: 1050}).Is(err) {
		// * return if error not caused by table already created
		return fmt.Errorf("Failed to create url table: %v", err)
	}

	return nil
}

// Implmentation of URL insert operation
func (dbHandler *MySQLURLDBHandler) Insert(entry models.URLEntry) error {
	newEntry := schemas.MySQLURLEntry{
		ID:       entry.ID,
		LongURL:  entry.LongURL,
		ShortURL: entry.ShortURL,
	}
	// TODO: refactor the sql staement to other place
	sqlStatement := `
INSERT INTO urlmappings (shortURL, longURL)
VALUES (:shortURL, :longURL)`
	_, err := dbHandler.conn.NamedExec(sqlStatement, newEntry)
	if err != nil {
		dbHandler.Logger.Log(err.Error())
		return err
	}
	return nil
}

func (dbHandler *MySQLURLDBHandler) Query(query models.URLEntry) (models.URLEntry, error) {
	queryModel := schemas.MySQLURLEntry{
		ID:       query.ID,
		ShortURL: query.ShortURL,
		LongURL:  query.LongURL,
	}

	filterString := dbutil.FilterString(queryModel)

	queryString := dbutil.SelectFields("urlmappings", queryModel, filterString)

	// * Execute query
	results, err := dbHandler.conn.Queryx(queryString)
	defer results.Close()
	if err != nil {
		dbHandler.Logger.Log(err.Error())
		return models.URLEntry{}, err
	}

	// * If query success
	if results.Next() {
		var entry schemas.MySQLURLEntry
		results.StructScan(&entry)
		return models.URLEntry{
			ID:       entry.ID,
			ShortURL: entry.ShortURL,
			LongURL:  entry.LongURL,
		}, nil
	}
	return models.URLEntry{}, nil
}

func (dbHandler *MySQLURLDBHandler) Close() error {
	return dbHandler.conn.Close()
}

func (dbHandler *MySQLURLDBHandler) loadDBConfig(config DBConfig) error {
	// * Connect to DB
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := sqlx.Open(config.Driver, dbSource)
	if err != nil {
		return fmt.Errorf("Fail to load db config: %v", err)
	}
	dbHandler.conn = db
	return nil
}

func (dbHandler *MySQLURLDBHandler) loadPoolConfig(config PoolConfig) {
	dbHandler.conn.SetMaxIdleConns(config.MaxIdleConns)
	dbHandler.conn.SetMaxOpenConns(config.MaxOpenConns)
	dbHandler.conn.SetConnMaxIdleTime(time.Duration(config.ConnMaxIdleTime))
	dbHandler.conn.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime))
}
