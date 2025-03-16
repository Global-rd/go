package jsonService

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	booksModel "restApiClient/models"
)

func OpenJsonFile(books *booksModel.BookModel, fileName string) {
	data, err := os.Open("resources/books.json")
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
	}

	defer data.Close()

	booksData, err := io.ReadAll(data)
	if err != nil {
		fmt.Errorf("error reading file: %v", err)
	}

	var booksTemp []booksModel.Book

	err = json.Unmarshal(booksData, &booksTemp)
	if err != nil {
		fmt.Errorf("error unmarshalling json: %v", err)
	}

	books.Books = booksTemp
}
