package schemas

type CreateURLResponse struct {
	LongURL  string `json:"longURL"`
	ShortURL string `json:"shortURL"`
}

type GetURLResponse struct {
	LongURL  string `json:"longURL"`
	ShortURL string `json:"shortURL"`
}
