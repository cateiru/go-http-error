package main

import (
	"errors"
	"fmt"

	"github.com/cateiru/go-http-error/httperror"
	"github.com/cateiru/go-http-error/httperror/status"
)

func main() {
	// Create 404 notfound error
	err := status.NewNotFoundError(errors.New("error message")).Wrap()

	fmt.Println(err.Error())

	// Create 404 error, and add filename and line.
	err = status.NewNotFoundError(errors.New("error message")).Caller("example.go", 17).AddCode(0).Wrap()

	fmt.Println(err.Error())

	if castedErr, ok := httperror.CastHTTPError(err); ok {
		fmt.Println(castedErr.FileName)
		fmt.Println(castedErr.Line)
		fmt.Println(castedErr.StatusCode)
		fmt.Println(castedErr.Code)
		fmt.Println(castedErr.Unwrap().Error())
	}

	// Create custom http status error
	err = httperror.NewError(301, errors.New("error message")).Wrap()

	fmt.Println(err.Error())
}
