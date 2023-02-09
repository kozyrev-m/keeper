package datamodel

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// LoginPassword contains login password pair.
type LoginPassword struct {
	BasePart
	Login    string
	Password string
}

// Validate validates login-password pair.
func (lp *LoginPassword) Validate() error {
	return validation.ValidateStruct(
		lp,
		validation.Field(&lp.Login, validation.Required, validation.Length(5, 20), is.Alphanumeric),
		validation.Field(&lp.Password, validation.Required, validation.Length(6, 20)),
	)
}

// Encrypt encrypts content.
func (lp *LoginPassword) Encrypt() error {

	value := strings.Join(
		[]string{lp.Login, lp.Password},
		separator,
	)

	enc, err := encrypt(password, value)
	if err != nil {
		return err
	}

	lp.SetEncryptedContent(enc)

	return nil
}

// Decrypt decrypts content.
func (lp *LoginPassword) Decrypt(enc string) error {
	value, err := decrypt(password, enc)
	if err != nil {
		return err
	}

	loginpassword := strings.Split(value, separator)

	lp.Login = loginpassword[0]
	lp.Password = loginpassword[1]

	return nil
}
