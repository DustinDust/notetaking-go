package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	conn *sqlx.DB
}

func (d *Database) StartDBConnection(DB_URI string) {
	var err error
	d.conn, err = sqlx.Connect("postgres", DB_URI)
	if err != nil {
		log.Fatalln(err)
	}
	err = d.conn.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func (d *Database) Seeds() {
	// TODO: implement this
}
