package services

import (
	"errors"
	"fmt"
	"urlshortener/domain"
)

//TODO: refactor error handling code snippets to a package

// Class that doing CRUD operation to url entries
type URLEntryInteractor struct {
	URLRepository domain.URLRepository
	HashGenerator domain.HashGenerator
}

// Find the original url that corresponding to.
func (interactor *URLEntryInteractor) GetURL(shortURL string) (string, error) {
	// * check shortURL length is legal
	if len(shortURL) > 7 || len(shortURL) < 1 {
		return "", errors.New("Expected url shorter than 7 characters")
	}

	queryEntry := domain.URLEntry{ShortURL: shortURL}
	entry, err := interactor.URLRepository.Get(queryEntry)
	if err != nil {
		return "", fmt.Errorf("Error occurred in url repository: %v", err)
	}

	// * check if query nothing
	if entry == *new(domain.URLEntry) {
		errMsg := fmt.Sprintf("Expected entry of shortURL:{%s} but get nothing", shortURL)
		return "", errors.New(errMsg)
	}

	// * check shortURL is same as entry
	if entry.ShortURL != shortURL {
		errMsg := fmt.Sprintf("Expected entry of shortURL:{%s} but get {%s}", shortURL, entry.ShortURL)
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
		return "", fmt.Errorf("Error occurred in url repository: %v", err)
	}

	// * check if entry is already exist in repository
	if entry != *new(domain.URLEntry) {
		return entry.ShortURL, nil
		// * create new entry/shorten if entry is not exist
	} else {
		hash := interactor.HashGenerator.GenerateHash()

		// * check url is llegal
		if len(hash) > 7 {
			return "", errors.New("Expected url shorter than 7 characters, check status of url shortener")
		}

		entry := domain.URLEntry{ShortURL: hash, LongURL: longURL}
		err = interactor.URLRepository.Store(entry)
		if err != nil {
			return "", fmt.Errorf("Error occurred when store url to repository: %v", err)
		}

		return hash, nil
	}
}
