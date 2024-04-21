package repositories

import "shortener/models"
import "shortener/domain"

// DB interface between framework and interface layer
type URLDBHandler interface {
	Insert(entry models.URLEntry) error
	Query(query models.URLEntry) (models.URLEntry, error)
}

// class that act as a interface between application and database
type URLRepository struct {
	dbHandler URLDBHandler
	logger    Logger
}

// store url entry into database
func (repo *URLRepository) Store(urlEntry domain.URLEntry) error {
	entry := models.URLEntry{
		ShortURL: urlEntry.ShortURL,
		LongURL:  urlEntry.LongURL,
	}

	err := repo.dbHandler.Insert(entry)
	if err != nil {
		repo.logger.Log(err.Error())
		return err
	}
	return nil
}

// get url entry from database
func (repo *URLRepository) Get(query domain.URLEntry) (domain.URLEntry, error) {
	queryModel := models.URLEntry{
		ShortURL: query.ShortURL,
		LongURL:  query.LongURL,
	}

	result, err := repo.dbHandler.Query(queryModel)
	if err != nil {
		repo.logger.Log(err.Error())
		return domain.URLEntry{}, err
	}

	entry := domain.URLEntry{
		ShortURL: result.ShortURL,
		LongURL:  result.LongURL,
	}
	return entry, nil
}
