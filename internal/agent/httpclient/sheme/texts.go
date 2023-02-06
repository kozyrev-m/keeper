package sheme

import (
	"encoding/json"
	"fmt"

	"github.com/kozyrev-m/keeper/internal/agent/model"
)

// ResponseTexts - response with texts.
type ResponseTexts struct {
	Texts []model.Text
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *ResponseTexts) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return fmt.Errorf("no bytes to unmarshal")
	}

	// See if we can guess based on the first character
	switch b[0] {
	case '{':
		return t.unmarshalSingle(b)
	case '[':
		return t.unmarshalMany(b)
	}

	// TODO: Figure out what do we do here
	return nil
}

func (t *ResponseTexts) unmarshalSingle(b []byte) error {
	var txt model.Text
	err := json.Unmarshal(b, &txt)
	if err != nil {
		return err
	}
	t.Texts = []model.Text{txt}
	return nil
}

func (t *ResponseTexts) unmarshalMany(b []byte) error {
	var texts []model.Text
	err := json.Unmarshal(b, &texts)
	if err != nil {
		return err
	}
	t.Texts = texts
	return nil
}
