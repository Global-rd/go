package task

type Task struct {
	ID          string
	Name        string
	Description string
}

func (t *Task) SetName(name string) {
	t.Name = name
}
