package file

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Error = errors.New("file error: %v")
)

// File represents a file to be stored
type File struct {
	Name string
	Url  string
	Data io.ReadCloser
}

// Store saves a file to the dir specified,
// creating the dir if it doesn't exist
func (f *File) Store(dir string) error {
	defer f.Data.Close()

	path := dir + f.Name
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf(Error.Error(), err)
	}
	defer file.Close()

	log.Printf("Storing file: %s", path)
	_, err = io.Copy(file, f.Data)
	if err != nil {
		return fmt.Errorf(Error.Error(), err)
	}

	return nil
}

func MkDir(dir string) error {
	// Delete dir files if exist
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf(Error.Error(), err)
	}

	// Create dir
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return fmt.Errorf(Error.Error(), err)
	}

	return nil
}
