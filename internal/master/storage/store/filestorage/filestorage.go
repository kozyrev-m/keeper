package filestorage

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

const (
	Dir = "/tmp/filestorage"
)

// Create creates some file.
func CreateFile(ownerid int, fname string, file multipart.File) error {
	dir := fmt.Sprintf("%s/%d", Dir, ownerid)
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		errIn := os.MkdirAll(dir, os.ModePerm)
		if errIn != nil {
			return err
		}
	}

	// Create file
	dst, err := os.Create(fmt.Sprintf("%s/%s", dir, fname))
	if err != nil {
		return err
	}

	defer func() {
		if err := dst.Close(); err != nil {
			log.Println(err)
		}
	}()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	return nil
}

func DeleteFile(filepath string) error {
	return os.Remove(filepath)
}
