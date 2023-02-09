package sheme

import (
	"encoding/json"
	"fmt"

	"github.com/kozyrev-m/keeper/internal/agent/model"
)

// ResponseCards - response with bank cards.
type ResponseCards struct {
	Cards []model.BankCard
}

// UnmarshalJSON implements json.Unmarshaler.
func (rc *ResponseCards) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return fmt.Errorf("no bytes to unmarshal")
	}

	// See if we can guess based on the first character
	switch b[0] {
	case '{':
		return rc.unmarshalSingle(b)
	case '[':
		return rc.unmarshalMany(b)
	}

	// TODO: Figure out what do we do here
	return nil
}

func (rc *ResponseCards) unmarshalSingle(b []byte) error {
	var bc model.BankCard
	err := json.Unmarshal(b, &bc)
	if err != nil {
		return err
	}
	rc.Cards = []model.BankCard{bc}
	return nil
}

func (rc *ResponseCards) unmarshalMany(b []byte) error {
	var cards []model.BankCard
	err := json.Unmarshal(b, &cards)
	if err != nil {
		return err
	}
	rc.Cards = cards
	return nil
}
