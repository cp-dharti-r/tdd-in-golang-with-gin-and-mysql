package main

import "github.com/jmoiron/sqlx"

func TestDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:password@/db-name")
	return db, err
}
