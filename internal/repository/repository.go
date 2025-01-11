package repository

import "database/sql"

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func Connect() *sql.DB {
	return nil
}

func TestConnect() *sql.DB {
	return nil
}
