package db

import "testing"

var _ DB = (*DBMock)(nil)

type DBMock struct {
	changed int
	err     error
}

func (d DBMock) Insert(value string) (int, error) {
	return d.changed, d.err
}

func TestCreateUser(t *testing.T) {
	db := DBMock{1, nil}

	user := User{
		db: db,
	}

	name := "Peter"

	err := user.Create(name)

	if err != nil {
		t.Errorf("CreateUser(%s): expected error occured: %s", name, err.Error())
	}
}

func TestCreateUserFail(t *testing.T) {
	db := DBMock{0, nil}

	user := User{
		db: db,
	}

	name := "Peter"

	err := user.Create(name)

	if err == nil {
		t.Errorf("CreateUser(%s): error expected, but it is nil", name)
	}
}
