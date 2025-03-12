package db

import (
	"encoding/json"
	"fmt"
	"os"
)

const DbFile = "bookdb.json"

var Cache map[string]Book

func init() {
	Cache = make(map[string]Book)
	file, err := os.Open(DbFile)
	if err != nil {
		panic(fmt.Sprintf("Initialization failed: Error opening file %s due to error: %v", DbFile, err))
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Cache)
	if err != nil {
		panic(fmt.Sprintf("Initialization failed: Error decoding file %s due to error: %v", DbFile, err))
	}
}

func GetBook(id string) (Book, error) {
	book, ok := Cache[id]
	if !ok {
		return Book{}, fmt.Errorf("book with id %s not found", id)
	}
	return book, nil
}
