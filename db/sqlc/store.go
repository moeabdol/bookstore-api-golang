package db

import "database/sql"

// Store struct
type Store struct {
	db *sql.DB
	*Queries
}

// NewStore function to create a new database store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
