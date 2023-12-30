package file

import (
	"bufio"
	"github.com/habibeh92/request-counter/internal/repository"
	"os"
	"strings"
)

type request struct {
	file *os.File
}

// New get the new instance of request repository
func New(file *os.File) repository.Request {
	return &request{file: file}
}

// Add append a string to the file
func (r *request) Add(data string) error {
	_, err := r.file.WriteString(data + "\n")

	return err
}

// All get a list of all lines of the file
func (r *request) All() ([]string, error) {
	var data []string
	scanner := bufio.NewScanner(r.file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, scanner.Err()
}

// Truncate clear the file
func (r *request) Truncate() error {
	return r.file.Truncate(0)
}

// Sync replace the list of data with the old ones
func (r *request) Sync(data []string) error {
	err := r.Truncate()
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	return r.Add(strings.Join(data, "\n"))
}
