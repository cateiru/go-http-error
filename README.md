# Go HTTP Error

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cateiru/go-http-error?style=flat-square) [![Go Reference](https://pkg.go.dev/badge/github.com/cateiru/go-http-error.svg)](https://pkg.go.dev/github.com/cateiru/go-http-error)

Wraps the normal error and provides an error that is easy to use with net/http.

## Install

```bash
go get -u github.com/cateiru/go-http-error
```

## Usage

```go
import (
    "github.com/cateiru/go-http-error/httperror"
    "github.com/cateiru/go-http-error/httperror/status"
)

// Create 404 notfound error
err := status.NewNotFoundError(errors.New("error message"))

// Create 404 error, and add custom code, filename and line.
err = status.NewNotFoundError(errors.New("error message")).Caller().AddCode(10)

// Cast to type HTTPError
if castedErr, ok := httperror.CastHTTPError(err); ok {
    fileName := castedErr.FileName
    line := castedErr.Line
    statusCode := castedErr.StatusCode
    code := castedErr.Code

    unwrapErr := castedErr.Unwrap()
    errMessage  := unwrapErr.Error()
}

// Create custom http status error
err = httperror.NewError(301, errors.New("error message"))
```

### Example

```bash
go run ./example/example.go
```

### Support HTTP Status

- See [client_error.go](./httperror/client_error.go) for HTTP status codes in the 400s.
- See [server_error.go](./httperror/server_error.go) for HTTP status codes in the 500s.

## LICENSE

[MIT](./LICENSE)
