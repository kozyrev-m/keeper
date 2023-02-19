package datamodel

import (
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// BankCard contains bank card data.
type BankCard struct {
	BasePart
	PAN       string `json:"pan"` // PAN (primary account number)
	CVV       string `json:"cvv"` // CVV/CVC (Card Verification Value/Code)
	ValidThru string `json:"valid_thru"`
	Name      string `json:"name"`
}

// Validate validates bank card data.
func (bc *BankCard) Validate() error {
	return validation.ValidateStruct(
		bc,
		validation.Field(&bc.PAN, validation.Required, validation.Length(16, 19), is.Digit),
		validation.Field(&bc.CVV, validation.Required, validation.Length(3, 3), is.Digit),
		validation.Field(&bc.ValidThru, validation.Required, validation.Date("02/06")),
		validation.Field(&bc.Name, validation.Required, validation.Length(2, 100), is.UpperCase),
	)
}

// Encrypt encrypts content.
func (bc *BankCard) Encrypt() error {
	value := fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		bc.PAN,
		separator,
		bc.CVV,
		separator,
		bc.ValidThru,
		separator,
		bc.Name,
	)

	enc, err := encrypt(password, value)
	if err != nil {
		return err
	}

	bc.SetEncryptedContent(enc)

	return nil
}

// Decrypt decrypts content.
func (bc *BankCard) Decrypt(enc string) error {
	value, err := decrypt(password, enc)
	if err != nil {
		return err
	}

	bankcard := strings.Split(value, separator)

	bc.PAN = bankcard[0]
	bc.CVV = bankcard[1]
	bc.ValidThru = bankcard[2]
	bc.Name = bankcard[3]

	return nil
}
