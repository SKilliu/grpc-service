package db

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

// QInterface represent interface of the access to the user db.
type QInterface interface {
	DBX() *dbx.DB
}

// DB wraps dbx interface.
type DB struct {
	db *dbx.DB
}

// DBX returns db client.
func (d DB) DBX() *dbx.DB {
	return d.db
}

// New connection opening.
func New(link string) (QInterface, error) {
	db, err := dbx.Open("postgres", link)
	return &DB{db: db}, err
}
