package models

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Models struct {
	db *sqlx.DB
}

func (m *Models) StartDBConnection(dbUrl string) {
	var err error
	m.db, err = sqlx.Connect("sqlite3", dbUrl)
	if err != nil {
		log.Fatalln(err)
	}
	err = m.db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}
