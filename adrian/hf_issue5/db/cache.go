package db

import (
	"encoding/json"
	"fmt"
	"os"
)

const DbFile = "bookdb.json"

var Cache map[string]Book

func LoadCache() error {
	Cache = make(map[string]Book)
	file, err := os.Open(DbFile)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Cache)
	if err != nil {
		return err
	}
	return nil
}

func GetBook(id string) (Book, error) {
	book, ok := Cache[id]
	if !ok {
		return Book{}, fmt.Errorf("book with id %s not found", id)
	}
	return book, nil
}
