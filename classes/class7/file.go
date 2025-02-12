package main

type File struct {
	Name string
}

func (f File) GetName() string {
	return f.Name
}
