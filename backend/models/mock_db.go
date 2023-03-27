package models

import (
	"database/sql"
	"errors"
)

type MockDB struct{}

func (m *MockDB) Connection() *sql.DB {
	return nil
}

func (m *MockDB) CreateUser(user User) error {
	if user.Name == "fail" {
		return errors.New("error creating user")
	}

	return nil
}

func (m *MockDB) GetUser(name string) (User, error) {
	var user User

	mockedUser := User{
		Name:     "testuser",
		Linkedin: "https://www.linkedin.com/in/testuser",
		Github:   "https://github.com/testuser",
	}

	if name == mockedUser.Name {
		return mockedUser, nil
	}

	if name == "" {
		return user, errors.New("error getting data")
	}

	return user, nil
}

func (m *MockDB) CheckUser(name string) error {
	if name == "newtestuser" || name == "fail" {
		return sql.ErrNoRows
	}
	return nil
}
