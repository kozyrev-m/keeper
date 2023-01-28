package datamodel

// Text contains text.
type Text struct {
	Value string
}

// Encrypt encrypts content.
func (t *Text) Encrypt() (string, error) {
	return "", nil
}

// Decrypt decrypts content.
func (t *Text) Decrypt(enc string) error {
	return nil
}
