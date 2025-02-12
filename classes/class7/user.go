package main

var entity Entity = (*User)(nil)

// type Address struct {
//
//}

// func (a Address) GetAddress() (interface {
// 	GetName() string
// }, error) {
// 	return User{}, nil
// }

type User struct {
	Base
	Name string
	// Address interface {
	// 	GetAddress() (interface {
	// 		GetName() string
	// 	}, error)
	// }
}

func (u User) GetName() string {
	return u.Name
}
