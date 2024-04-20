package services

import (
	"errors"
	"fmt"
	"shortener/domain"
)

//TODO: refactor error handling code snippets to a package

// Interface that can generate a short url from a long url
type Shortener interface {
	GenerateShortURL(longURL string) (string, error)
}

// Class that doing CRUD operation to url entries
type URLEntryInteractor struct {
	URLRepository domain.URLRepository
	Logger        Logger
	URLShortener  Shortener
}

// Find the original url that corresponding to.
func (interactor *URLEntryInteractor) GetURL(shortURL string) (string, error) {
	entry, err := interactor.URLRepository.Get(shortURL)
	if err != nil {
		interactor.Logger.Log(err.Error())
		return "", err
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
	shortURL, err := interactor.URLShortener.GenerateShortURL(longURL)
	if err != nil {
		interactor.Logger.Log(err.Error())
		return "", err
	}

	// * check url is llegal
	if len(shortURL) > 7 {
		interactor.Logger.Log("Expected url shorter than 7 characters, check status of url shortener")
	}

	entry := domain.URLEntry{
		ShortURL: shortURL,
		LongURL:  longURL,
	}
	err = interactor.URLRepository.Store(entry)
	if err != nil {
		interactor.Logger.Log(err.Error())
		return "", err
	}

	return shortURL, nil
}
