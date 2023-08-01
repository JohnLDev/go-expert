package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewConnection() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	// migration
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS dolar_price ( id text primary key, code text, codeIn text, name text, high text, low text, varBid text, pctChange text, bid text, ask text, timestamp text, createDate text)`)
	if err != nil {
		panic(err)
	}

	return db
}
