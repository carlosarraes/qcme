package models

import (
	"database/sql"
	"errors"
)

type Postgress struct {
	*sql.DB
}

type User struct {
	Name     string `json:"name"`
	Linkedin string `json:"linkedIn"`
	Github   string `json:"gitHub"`
}

func (m *Postgress) Connection() *sql.DB {
	return m.DB
}

func (m *Postgress) CreateUser(user User) error {
	data, err := m.Exec("INSERT INTO data.qrcode (name, linkedin, github) VALUES ($1, $2, $3)", user.Name, user.Linkedin, user.Github)
	if err != nil {
		return err
	}

	status, _ := data.RowsAffected()
	if status == 0 {
		myErr := errors.New("error inserting data")
		return myErr
	}

	return nil
}

func (m *Postgress) GetUser(name string) (User, error) {
	var user User
	err := m.QueryRow("SELECT name, linkedin, github FROM data.qrcode WHERE name = $1", name).Scan(&user.Name, &user.Linkedin, &user.Github)
	if err != nil {
		return user, err
	}

	return user, nil
}
