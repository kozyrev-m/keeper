// Package sqlstore provides sql store implementation.
package sqlstore

import (
	"database/sql"

	"github.com/kozyrev-m/keeper/internal/master/model"
	"github.com/kozyrev-m/keeper/internal/master/storage/store"
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

// CreateUser creates user.
func (s *Store) CreateUser(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return s.db.QueryRow(
		"INSERT INTO users (login, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Login,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

// FindUserByLogin finds user by login.
func (s *Store) FindUserByLogin(login string) (*model.User, error) {
	u := &model.User{}

	if err := s.db.QueryRow(
		"SELECT id, login, encrypted_password FROM users WHERE login = $1",
		login,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// FindUserByID finds user by login.
func (s *Store) FindUserByID(id int) (*model.User, error) {
	u := &model.User{}

	if err := s.db.QueryRow(
		"SELECT id, login, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Login,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}