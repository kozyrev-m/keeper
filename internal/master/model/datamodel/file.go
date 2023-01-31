package datamodel

import (
	"strings"
)

// File contains file information.
type File struct {
	ID       int    `json:"id"`
	Metadata string `json:"metadata"`
	Filename string `json:"filename"`
	Filepath string `json:"-"`
}

// Name gives file name.
func (f *File) Name() {
	els := strings.Split(f.Filepath, "/")
	
	f.Filename = els[len(els) - 1]
}
