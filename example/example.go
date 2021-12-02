package main

import (
	"errors"
	"fmt"

	"github.com/cateiru/go-http-error/httperror"
)

func main() {
	// Create 404 notfound error
	err := httperror.NewNotFoundError(errors.New("error message")).Wrap()

	fmt.Println(err.Error())

	// Create 404 error, and add filename and line.
	err = httperror.NewNotFoundError(errors.New("error message")).Caller("example.go", 17).Wrap()

	fmt.Println(err.Error())

	// Create custom http status error
	err = httperror.NewError(301, errors.New("error message")).Wrap()

	fmt.Println(err.Error())
}
