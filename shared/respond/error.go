package respond

import "fmt"

var (
	// NotFoundErr is a 404 error
	NotFoundErr = RequestError{404, "Not found"}
	// ForbiddenErr is a 403 error
	ForbiddenErr = RequestError{403, "Forbidden"}
)

// Error is an error to return in an http response
type Error interface {
	StatusCode() int
	Error() string
}

// RequestError is an error to return in an http response
type RequestError struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

func (e RequestError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Status, e.Message)
}

// StatusCode returns the status code
func (e RequestError) StatusCode() int {
	return e.Status
}

// NewRequestError creates a new error
func NewRequestError(status int, message string) error {
	return &RequestError{status, message}
}
