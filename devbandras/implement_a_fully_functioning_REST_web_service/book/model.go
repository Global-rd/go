package book

// Egy könyvet leíró struktúra
type Book struct {
	ID            int64   `db:"id" json:"id"`
	Title         string  `db:"title" json:"title"`
	Author        string  `db:"author" json:"author"`
	PublishedYear int     `db:"published_year" json:"published_year"`
	Genre         string  `db:"genre" json:"genre"`
	Price         float32 `db:"price" json:"price"`
}

// Válaszüzenetet reprezentáló struktúra
type BookResponseMessage struct {
	Message string `json:"message"`
}
