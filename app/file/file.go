package file

import (
	"errors"
	"fmt"
	"io"
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

// Store saves a file to the dir specified, creating the dir if it doesn't exist
func (f *File) Store(dir string) error {
	if err := mkDir(dir); err != nil {
		return err
	}

	path := dir + f.Name
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf(Error.Error(), err)
	}
	defer file.Close()
	defer f.Data.Close()

	_, err = io.Copy(file, f.Data)
	if err != nil {
		return fmt.Errorf(Error.Error(), err)
	}

	return nil
}

func mkDir(dir string) error {
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return fmt.Errorf(Error.Error(), err)
		}
	}

	return nil
}
