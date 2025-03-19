package book

type Book struct {
	ID            int64   `json:"id"`
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	PublishedYear int8    `json:"published_year"`
	Genre         string  `json:"genre"`
	Price         float32 `json:"price"`
}
