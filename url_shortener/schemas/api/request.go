package schemas

type CreateURLRequest struct {
	ID       int64  `json:"id"`
	LongURL  string `json:"longURL"`
	ShortURL string `json:"shortURL"`
}
