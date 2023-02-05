package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// BankCard contains bank card data.
type BankCard struct {
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