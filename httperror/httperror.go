package httperror

import (
	"fmt"
)

type HTTPError struct {
	StatusCode int

	FileName string
	Line     int

	Err error
}

// Override error.Error().
//
// It can be called in the same way as the normal err.Error().
//
// Example:
//	errMessage := err.Error()
func (c *HTTPError) Error() string {
	if len(c.FileName) != 0 {
		return fmt.Sprintf("\"%s\", line %d, %d:: %s", c.FileName, c.Line, c.StatusCode, c.Err.Error())
	}
	return fmt.Sprintf("%d:: %s", c.StatusCode, c.Err.Error())
}

// Unwrap the error.
//
// Example:
//	err := err.Unwrap()
func (c *HTTPError) Unwrap() error {
	return c.Err
}

// Create an HTTPError.
//
// Example:
//	err := NewError(400, errors.New("Bad Request"))
func NewError(statusCode int, err error) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		FileName:   "",
		Line:       0,
		Err:        err,
	}
}

// Specifies the file name and line number when an error occurs.
//
// This method is optional.
//
// Example:
//	err := NewNotFoundError(err).Caller("example.go", 10).Wrap()
func (c *HTTPError) Caller(fileName string, line int) *HTTPError {
	c.FileName = fileName
	c.Line = line

	return c
}

// Convert to error type.
//
// Example:
//	err := NewNotFoundError(err).Wrap()
func (c *HTTPError) Wrap() error {
	return c
}

// Casts the error to type HTTPError.
//
// If it cannot be cast, it returns (nil, false).
//
// Example:
//	if httpError, ok := CastHTTPError(err); ok {
//		...
//	}
func CastHTTPError(err error) (*HTTPError, bool) {
	if err, ok := err.(*HTTPError); ok {
		return err, true
	}
	return nil, false
}
