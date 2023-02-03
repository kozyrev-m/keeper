package datamodel

import (
	"fmt"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// BankCard contains bank card data.
type BankCard struct {
	BasePart
	PAN       uint64 `json:"pan"` // PAN (primary account number)
	CVV       uint8  `json:"cvv"` // CVV/CVC (Card Verification Value/Code)
	ValidThru string `json:"valid_thru"`
	Name      string `json:"name"`
}

// Validate validates bank card data.
func (bc *BankCard) Validate() error {
	return validation.ValidateStruct(
		bc,
		validation.Field(&bc.PAN, validation.Required),
		validation.Field(&bc.CVV, validation.Required),
		validation.Field(&bc.ValidThru, validation.Required, validation.Date("02/06")),
		validation.Field(&bc.Name, validation.Required, validation.Length(2, 100), is.Alpha),
	)
}

// Encrypt encrypts content.
func (bc *BankCard) Encrypt() error {
	value := fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		strconv.FormatUint(bc.PAN, 10),
		separator,
		strconv.Itoa(int(bc.CVV)),
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

	pan, err := strconv.ParseUint(bankcard[0], 10, 0)
	if err != nil {
		return err
	}
	bc.PAN = pan

	cvv, err := strconv.ParseInt(bankcard[1], 10, 8)
	if err != nil {
		return err
	}
	bc.CVV = uint8(cvv)

	bc.ValidThru = bankcard[2]
	bc.Name = bankcard[3]

	return nil
}
