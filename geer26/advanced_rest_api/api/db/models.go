package db

type Book struct {
	Id           string  `json:"id"`
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Published    int     `json:"published"`
	Introduction string  `json:"introduction"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
}
