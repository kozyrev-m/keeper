// Package sqlstore provides sql store implementation.
package sqlstore

import (
	"database/sql"
)

// Store contains sql storage implementation.
type Store struct {
	db *sql.DB
}

// New creates sql storage.
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
