// Package model contains models.
package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

// User contains information about user.
type User struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Password 		  string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

// Validate checks if the fields are correct.
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Login, validation.Required),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 100)),
	)
}

// BeforeCreate prepares the model to be added to the store.
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)

		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

// encryptString encrypts string to hash.
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	
	if err != nil {
		return "", err
	}

	return string(b), nil
}