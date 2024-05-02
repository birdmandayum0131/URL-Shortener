package schemas

type MySQLURLEntry struct {
	// * json tag for json row binding
	// * db tag for insert statement mapping
	ID       int64  `db:"id"`
	LongURL  string `db:"longURL"`
	ShortURL string `db:"shortURL"`
}
