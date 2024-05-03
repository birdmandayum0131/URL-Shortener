package services

import (
	"errors"
	"fmt"
	"urlshortener/domain"
	"logger"
)

//TODO: refactor error handling code snippets to a package

// Class that doing CRUD operation to url entries
type URLEntryInteractor struct {
	URLRepository domain.URLRepository
	Logger        logger.Logger
	HashGenerator domain.HashGenerator
}

// Find the original url that corresponding to.
func (interactor *URLEntryInteractor) GetURL(shortURL string) (string, error) {
	queryEntry := domain.URLEntry{ShortURL: shortURL}
	entry, err := interactor.URLRepository.Get(queryEntry)
	if err != nil {
		interactor.Logger.Log(err.Error())
		return "", err
	}

	// * check if query nothing
	if entry == *new(domain.URLEntry) {
		errMsg := fmt.Sprintf("Expected entry of shortURL:{%s} but get nothing", shortURL)
		interactor.Logger.Log(errMsg)
		return "", errors.New(errMsg)
	}

	// * check shortURL is same as entry
	if entry.ShortURL != shortURL {
		errMsg := fmt.Sprintf("Expected entry of shortURL:{%s} but get {%s}", shortURL, entry.ShortURL)
		interactor.Logger.Log(errMsg)
		return "", errors.New(errMsg)
	}

	return entry.LongURL, nil
}

// Generate and store a new url entry
func (interactor *URLEntryInteractor) CreateEntry(longURL string) (string, error) {
	queryEntry := domain.URLEntry{LongURL: longURL}
	// Check url entry is already exist in repository
	entry, err := interactor.URLRepository.Get(queryEntry)
	if err != nil {
		interactor.Logger.Log(err.Error())
		return "", err
	}

	// * check if entry is already exist in repository
	if entry != *new(domain.URLEntry) {
		msg := fmt.Sprintf("Entry of URL:{%s} already exist in repository", longURL)
		interactor.Logger.Log(msg)
		return entry.ShortURL, nil
	} else {
		hash := interactor.HashGenerator.GenerateHash()

		// * check url is llegal
		if len(hash) > 7 {
			interactor.Logger.Log("Expected url shorter than 7 characters, check status of url shortener")
		}

		entry := domain.URLEntry{ShortURL: hash, LongURL: longURL}
		err = interactor.URLRepository.Store(entry)
		if err != nil {
			interactor.Logger.Log(err.Error())
			return "", err
		}

		return hash, nil
	}
}
