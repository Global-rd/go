package database

import "fmt"

type Book struct {
	Id           int
	Title        string
	Author       string
	Published    int
	Introduction string
	Price        float32
	Stock        int
}

type Store struct {
	Books []Book
}

func DialStore() (*Store, error) {
	store := Store{}

	if err := store.LoadStore(); err != nil {
		return nil, fmt.Errorf("error at dialing database: %s", err.Error())
	}

	return &store, nil
}

func (s Store) LoadStore() error {
	return nil
}

func (s Store) FlushStore() error {
	return nil
}

func (s Store) FindOne(key string, value any) (Book, error) {
	book := Book{}
	return book, nil
}

func (s Store) FindMany(key string, value any) ([]Book, error) {
	shelf := []Book{}
	return shelf, nil
}

func (s Store) DeleteOne(key string, value any) error {
	return nil
}

func (s Store) DeleteMany(key string, value any) error {
	return nil
}

func (s Store) UpdateOne(key string, value any, new_content Book) error {
	return nil
}

func (s Store) CreateOne(book Book) error {
	return nil
}
