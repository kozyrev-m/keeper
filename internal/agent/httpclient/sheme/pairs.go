package sheme

import (
	"encoding/json"
	"fmt"

	"github.com/kozyrev-m/keeper/internal/agent/model"
)

// ResponseCards - response with login-password pairs.
type ResponsePairs struct {
	Pairs []model.Pair
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *ResponsePairs) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return fmt.Errorf("no bytes to unmarshal")
	}

	// See if we can guess based on the first character
	switch b[0] {
	case '{':
		return p.unmarshalSingle(b)
	case '[':
		return p.unmarshalMany(b)
	}

	// TODO: Figure out what do we do here
	return nil
}

func (p *ResponsePairs) unmarshalSingle(b []byte) error {
	var pr model.Pair
	err := json.Unmarshal(b, &pr)
	if err != nil {
		return err
	}
	p.Pairs = []model.Pair{pr}
	return nil
}

func (p *ResponsePairs) unmarshalMany(b []byte) error {
	var pairs []model.Pair
	err := json.Unmarshal(b, &pairs)
	if err != nil {
		return err
	}
	p.Pairs = pairs
	return nil
}
