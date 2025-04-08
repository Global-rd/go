package payment

type Payment struct {
	ID          string
	Name        string
	Description string
}

func (t *Payment) SetName(name string) {
	t.Name = name
}
