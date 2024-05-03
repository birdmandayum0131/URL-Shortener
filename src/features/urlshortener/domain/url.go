package domain

// Repository that store the entry of url mapping
type URLRepository interface {
	Store(entry URLEntry) error
	Get(query URLEntry) (URLEntry, error)
}

// Entry of url mapping task
type URLEntry struct {
	ShortURL string
	LongURL  string
}
