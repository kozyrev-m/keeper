package httpclient

import (
	"encoding/json"
	"fmt"
)

// responseError - scheme for response error.
type responseError struct {
	Error string `json:"error"`
}

// File - scheme for file response.
type File struct {
	ID    int  `json:"id"`
	Metadata string `json:"metadata"`
	Name  string `json:"filename"`
}

// responseFiles - response with files.
type ResponseFiles struct {
	Files []File
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
	var f File
	err := json.Unmarshal(b, &f)
	if err != nil {
		return err
	}
	rf.Files = []File{f}
	return nil
}

func (rf *ResponseFiles) unmarshalMany(b []byte) error {
	var files []File
	err := json.Unmarshal(b, &files)
	if err != nil {
		return err
	}
	rf.Files = files
	return nil
}
