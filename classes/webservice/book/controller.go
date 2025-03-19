package book

import "webservice/container"

type Controller struct {
	books map[string]Book
	cont  container.Container
}

func NewController(cont container.Container) Controller {
	return Controller{
		books: map[string]Book{
			"id1": {ID: "test id", Title: "test title"},
		},
		cont: cont,
	}
}

func (c Controller) Create(book Book) error {
	return nil
}

func (c Controller) GetByID(id string) (Book, error) {
	return c.books[id], nil
}

func (c Controller) Get() ([]Book, error) {
	results := make([]Book, len(c.books))

	return results, nil
}

func (c Controller) Update(id string) error {
	return nil
}

func (c Controller) Delete(id string) error {
	return nil
}
