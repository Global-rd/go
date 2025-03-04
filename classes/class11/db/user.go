package db

import "errors"

type PSQLDB struct{}

func (*PSQLDB) Insert(value string) (int, error) {
	return 0, nil
}

type DB interface {
	Insert(value string) (int, error)
}

type User struct {
	db DB
}

func (u User) Create(name string) error {
	changed, err := u.db.Insert(name)
	if err != nil {
		return err
	}

	if changed == 0 {
		return errors.New("this query did not change the table")
	}

	return nil
}
