package book

type Controller struct {
	books map[string]Book
}

func NewController() Controller {
	return Controller{
		books: map[string]Book{
			"id1": {ID: "test id", Title: "test title"},
		},
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

	for _, book := range c.books {
		results = append(results, book)
	}

	return results, nil
}

func (c Controller) Update(id string) error {
	return nil
}

func (c Controller) Delete(id string) error {
	return nil
}
