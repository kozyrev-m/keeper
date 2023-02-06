package model

// File - scheme for file response.
type File struct {
	ID       int    `json:"id"`
	Metadata string `json:"metadata"`
	Name     string `json:"filename"`
}
