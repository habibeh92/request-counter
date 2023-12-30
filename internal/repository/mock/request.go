// A simple mock of request repository

package mock

import (
	"github.com/habibeh92/request-counter/internal/repository"
)

type mockRequest struct {
}

// New get the mock instance of an implementation of repository.Request
func New() repository.Request {
	return &mockRequest{}
}

// Add provides a mock function with given fields: data
func (r *mockRequest) Add(data string) error {
	return nil
}

// All provides a mock function which returns empty slice and nil
func (r *mockRequest) All() ([]string, error) {
	return []string{}, nil
}

// Truncate provides a mock function which returns nil
func (r *mockRequest) Truncate() error {
	return nil
}

// Sync provides a mock function which returns nil
func (r *mockRequest) Sync(data []string) error {
	return nil
}
