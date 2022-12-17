package httperror

import (
	"errors"
	"fmt"
	"runtime"
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
//
//	errMessage := err.Error()
func (c *HTTPError) Error() string {
	if len(c.FileName) != 0 {
		return fmt.Sprintf("%s:%d, %d, %s", c.FileName, c.Line, c.StatusCode, c.Err.Error())
	}
	return fmt.Sprintf("%d, %s", c.StatusCode, c.Err.Error())
}

// Unwrap the error.
//
// Example:
//
//	err := err.Unwrap()
func (c *HTTPError) Unwrap() error {
	return c.Err
}

// Create an HTTPError.
//
// Example:
//
//	err := NewError(400, errors.New("Bad Request"))
func NewError(statusCode int, err error) *HTTPError {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return &HTTPError{
			StatusCode: statusCode,
			FileName:   "",
			Line:       0,
			Err:        err,
		}
	}

	return &HTTPError{
		StatusCode: statusCode,
		FileName:   file,
		Line:       line,
		Err:        err,
	}
}

// Create an HTTPError used error message
func NewStringError(statusCode int, message string) *HTTPError {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return &HTTPError{
			StatusCode: statusCode,
			FileName:   "",
			Line:       0,
			Err:        errors.New(message),
		}
	}

	return &HTTPError{
		StatusCode: statusCode,
		FileName:   file,
		Line:       line,
		Err:        errors.New(message),
	}
}

// Casts the error to type HTTPError.
//
// If it cannot be cast, it returns (nil, false).
//
// Example:
//
//	if httpError, ok := CastHTTPError(err); ok {
//		...
//	}
func CastHTTPError(err error) (*HTTPError, bool) {
	if err, ok := err.(*HTTPError); ok {
		return err, true
	}
	return nil, false
}
