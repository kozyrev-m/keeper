package sqlstore

import (
	"database/sql"

	"github.com/kozyrev-m/keeper/internal/master/model/usermodel"
	"github.com/kozyrev-m/keeper/internal/master/storage/store"
)

// CreateUser creates user.
func (s *Store) CreateUser(u *usermodel.User) error {
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
func (s *Store) FindUserByLogin(login string) (*usermodel.User, error) {
	u := &usermodel.User{}

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
func (s *Store) FindUserByID(id int) (*usermodel.User, error) {
	u := &usermodel.User{}

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