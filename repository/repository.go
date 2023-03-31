package repository

import "database/sql"

type Repo struct {
	db *sql.DB
}

type RepoInterface interface {
	BookRepo
}

// constructor function
func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db} // handle dependencies
}
