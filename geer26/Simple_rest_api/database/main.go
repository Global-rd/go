package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

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
	Books []Book `json:"books"`
}

var books *Store

func DialStore() (*Store, error) {
	if books == nil {
		log.Println("No instance, create one!")
		books = &Store{}
		if err := books.LoadStore(); err != nil {
			return nil, err
		}
		return books, nil
	}
	log.Println("Has instance, return it!")
	return books, nil
}

func (s *Store) LoadStore() error {
	file, err := os.ReadFile("database/db.json")
	if err != nil {
		return fmt.Errorf("error at opening database: %s", err.Error())
	}
	if err = json.Unmarshal(file, &s); err != nil {
		return fmt.Errorf("error at parsing database: %s", err.Error())
	}
	return nil
}

func (s Store) FlushStore() error {
	err := os.Remove("database/db.json")
	if err != nil {
		return fmt.Errorf("error at flushing db: %s", err.Error())
	}
	jsonString, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("error at encoding db: %s", err.Error())
	}
	os.WriteFile("database/db.json", jsonString, os.ModePerm)
	return nil
}

func (s Store) FindOne(key string, value any) (Book, error) {
	book := Book{}
	return book, nil
}

func (s Store) FindAll() ([]Book, error) {
	return s.Books, nil
}

func (s Store) DeleteOne(key string, value any) error {
	return nil
}

func (s Store) DeleteAll() error {
	return nil
}

func (s Store) UpdateOne(key string, value any, new_content Book) error {
	return nil
}

func (s Store) CreateOne(book Book) error {
	return nil
}
