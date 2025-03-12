package db

type Book struct {
	Author        string `json:"author"`
	Title         string `json:"title"`
	ISBN          string `json:"isbn"`
	Pages         int    `json:"pages"`
	Publishing    string `json:"publishing"`
	PublishedDate int    `json:"publishedDate"`
}
