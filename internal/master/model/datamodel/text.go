package datamodel

const (
	password = "secretkey"
)

// Text contains text.
type Text struct {
	Value string
}

// Encrypt encrypts content.
func (t *Text) Encrypt() (string, error) {
	return encrypt(password, t.Value)
}

// Decrypt decrypts content.
func (t *Text) Decrypt(enc string) error {
	value, err := decrypt(password, enc)
	if err != nil {
		return err
	}

	t.Value = value

	return nil
}
