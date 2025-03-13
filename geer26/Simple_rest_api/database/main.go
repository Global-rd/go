package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Id           string  `json:"Id"`
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Introduction string  `json:"introduction"`
	Price        float32 `json:"price"`
	Stock        int     `json:"stock"`
}

type Store struct {
	Books []Book `json:"books"`
}

var books *Store

func DialStore() (*Store, error) {
	if books == nil {
		books = &Store{}
		if err := books.LoadStore(); err != nil {
			return nil, err
		}
		return books, nil
	}
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

func (s Store) FindOne(key string, value interface{}) (Book, error) {
	for _, b := range s.Books {
		var temp map[string]interface{}
		temp_marshalled, _ := json.Marshal(b)
		json.Unmarshal(temp_marshalled, &temp)
		if temp[key] == value {
			return b, nil
		}
	}
	return Book{}, fmt.Errorf("no Entry found with %s:%s specification", key, value)
}

func (s Store) FilterAll(key string, value interface{}) ([]Book, error) {
	var retval []Book
	for _, b := range s.Books {
		var temp map[string]interface{}
		temp_marshalled, _ := json.Marshal(b)
		json.Unmarshal(temp_marshalled, &temp)
		if temp[key] == value {
			retval = append(retval, b)
		}
	}
	return retval, nil
}

func (s Store) FindAll() ([]Book, error) {
	return s.Books, nil
}

func (s Store) DeleteOne(key string, value any) error {
	for i, b := range s.Books {
		var temp map[string]interface{}
		temp_marshalled, _ := json.Marshal(b)
		json.Unmarshal(temp_marshalled, &temp)
		if temp[key] == value {
			copy(s.Books[i:], s.Books[i+1:])
			s.Books = s.Books[:len(s.Books)-1]
			err := s.FlushStore()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (s Store) DeleteAll() error {
	s.Books = []Book{}
	err := s.FlushStore()
	if err != nil {
		return err
	}
	return nil
}

func (s Store) UpdateOne(key string, value any, new_content Book) error {
	return nil
}

func (s Store) CreateOne(book Book) error {
	return nil
}
