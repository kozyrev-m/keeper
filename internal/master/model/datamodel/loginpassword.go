package datamodel

import "strings"

// LoginPassword contains login password pair.
type LoginPassword struct {
	BasePart
	Login    string
	Password string
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
