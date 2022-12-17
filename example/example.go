package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/cateiru/go-http-error/httperror"
)

func main() {
	err := httperror.NewError(http.StatusNotFound, errors.New("error message"))

	fmt.Println(err.Error())

	if castedErr, ok := httperror.CastHTTPError(err); ok {
		fmt.Println(castedErr.FileName)
		fmt.Println(castedErr.Line)
		fmt.Println(castedErr.StatusCode)
		fmt.Println(castedErr.Unwrap().Error())
	}

	// Create custom http status error
	err = httperror.NewError(301, errors.New("error message"))

	fmt.Println(err.Error())
}
