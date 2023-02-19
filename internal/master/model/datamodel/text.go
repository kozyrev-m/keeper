package datamodel

import validation "github.com/go-ozzo/ozzo-validation"

// Text contains text.
type Text struct {
	BasePart
	Value string `json:"text"` // Main part of the record 
}

// Validate validates text.
func (t *Text) Validate() error {
	return validation.ValidateStruct(
		t,
		validation.Field(&t.Value, validation.Required, validation.Length(1, 150)),
	)
}

// Encrypt encrypts content.
func (t *Text) Encrypt() error {
	enc, err := encrypt(password, t.Value)
	if err != nil {
		return err
	}

	t.SetEncryptedContent(enc)

	return nil
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
