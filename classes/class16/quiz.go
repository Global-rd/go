package main

type User struct {
	Name string
}

func (u *User) SetName(name string) {
	if validate(name) {
		u.Name = name
	}
}

func validate(value string) bool {
	// validation logic here...

	return true
}
