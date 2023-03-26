package models

import "database/sql"

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
