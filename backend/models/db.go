package models

import (
	"database/sql"
	"errors"
	"strings"
)

type Postgres struct {
	*sql.DB
}

type User struct {
	Name     string `json:"name"`
	Linkedin string `json:"linkedIn"`
	Github   string `json:"gitHub"`
}

func (m *Postgres) Connection() *sql.DB {
	return m.DB
}

func (m *Postgres) CreateUser(user User) error {
	data, err := m.Exec("INSERT INTO data.qrcode (name, linkedin, github) VALUES ($1, $2, $3)", strings.ToLower(user.Name), user.Linkedin, user.Github)
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

func (m *Postgres) GetUser(name string) (User, error) {
	var user User
	err := m.QueryRow("SELECT name, linkedin, github FROM data.qrcode WHERE name = $1", name).Scan(&user.Name, &user.Linkedin, &user.Github)
	if err != nil || err == sql.ErrNoRows {
		return user, err
	}

	return user, nil
}

func (m *Postgres) CheckUser(name string) error {
	var existingName string
	err := m.QueryRow("SELECT name FROM data.qrcode WHERE name = $1", name).Scan(&existingName)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}

	return nil
}
