package infrastructure

import (
	"fmt"
	"logger"
	"reflect"
	"strings"
	"urlshortener/interfaces/models"
	"urlshortener/interfaces/schemas"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLURLDBHandler struct {
	Conn   *sqlx.DB
	Logger logger.Logger
}

// Initialize mysql url database
func (dbHandler *MySQLURLDBHandler) Init() {
	// * create url table
	// TODO: refactor the sql statement to better write style
	var schema = `
	CREATE TABLE url_shortener.urlmappings (
		id BIGINT NOT NULL AUTO_INCREMENT,
		shortURL VARCHAR(64) NULL,
		longURL VARCHAR(2048) NULL,
		PRIMARY KEY (id))`

	_, err := dbHandler.Conn.Exec(schema)
	// * check if table is created
	if err != nil && !(&mysql.MySQLError{Number: 1050}).Is(err) {
		// * panic if error not caused by table already created
		panic(err.Error())
	}
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
	_, err := dbHandler.Conn.NamedExec(sqlStatement, newEntry)
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
	fields := reflect.ValueOf(queryModel)
	types := fields.Type()
	filterStrings := make([]string, 0)
	// * iterate all fields, if not empty, append to filterStrings
	for i := 0; i < fields.NumField(); i++ {
		if !fields.Field(i).IsZero() {
			var compareString string
			if types.Field(i).Type == reflect.TypeOf("") {
				compareString = "%s = '%s'"
			} else {
				compareString = "%s = %s"
			}
			filterStrings = append(filterStrings, fmt.Sprintf(compareString, types.Field(i).Tag.Get("db"), fields.Field(i)))
		}
	}
	// TODO: refactor the sql filter statement to other place
	filterString := strings.Join(filterStrings, " AND ")

	results, err := dbHandler.Conn.Queryx(fmt.Sprintf("SELECT id, shortURL, longURL FROM urlmappings WHERE %s", filterString))
	if err != nil {
		dbHandler.Logger.Log(err.Error())
		return models.URLEntry{}, err
	}

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
