package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "user=postgres password=T0xic11Red0 dbname=ffmaindb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
