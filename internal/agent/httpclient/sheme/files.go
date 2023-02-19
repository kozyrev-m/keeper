package sheme

import (
	"encoding/json"
	"fmt"

	"github.com/kozyrev-m/keeper/internal/agent/model"
)

// responseFiles - response with files.
type ResponseFiles struct {
	Files []model.File
}

// UnmarshalJSON implements json.Unmarshaler.
func (rf *ResponseFiles) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return fmt.Errorf("no bytes to unmarshal")
	}

	// See if we can guess based on the first character
	switch b[0] {
	case '{':
		return rf.unmarshalSingle(b)
	case '[':
		return rf.unmarshalMany(b)
	}

	// TODO: Figure out what do we do here
	return nil
}

func (rf *ResponseFiles) unmarshalSingle(b []byte) error {
	var f model.File
	err := json.Unmarshal(b, &f)
	if err != nil {
		return err
	}
	rf.Files = []model.File{f}
	return nil
}

func (rf *ResponseFiles) unmarshalMany(b []byte) error {
	var files []model.File
	err := json.Unmarshal(b, &files)
	if err != nil {
		return err
	}
	rf.Files = files
	return nil
}