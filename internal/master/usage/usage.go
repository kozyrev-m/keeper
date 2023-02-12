// Package usage contains any methods for project.
package usage

import (
	"io"
	"log"
	"os"
)

func ReadFromFile(path string) (string, error) {
	if len(path) == 0 {
		return "", nil
	}

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	
	defer func () {
		if errF := f.Close(); errF != nil {
			log.Fatal(errF)
		}
	} ()

	file, err := io.ReadAll(f)

	if err != nil {
		return "", err
	}

	fileContent := string(file)

	return fileContent, nil
}