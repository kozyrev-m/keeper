package datamodel

type Content interface {
	Encrypt() (string, error)
	Decrypt(string) error
}

// Text contains text.
type Text struct {
	Text string
}

func (t *Text) Encrypt() (string, error) {
	return "", nil
}

func (t *Text) Decrypt(enc string) error {
	return nil
}
