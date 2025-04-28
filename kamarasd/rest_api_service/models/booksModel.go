package booksModel

type Book struct {
	ID     int    `json:"Id"`
	Title  string `json:"Title"`
	Writer string `json:"Writer"`
	Genre  string `json:"Genre"`
	Date   string `json:"Date"`
	ISBN   string `json:"ISBN"`
}

type BookModel struct {
	Books []Book
}

func NewBookModel() *BookModel {
	return &BookModel{}
}
