# Go HTTP Error

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
err := status.NewNotFoundError(errors.New("error message")).Wrap()

// Create 404 error, and add filename and line.
fileName := "example.com"
line := 17
err = status.NewNotFoundError(errors.New("error message")).Caller(fileName, line).Wrap()

// Cast to type HTTPError
if castedErr, ok := httperror.CastHTTPError(err); ok {
    fileName := castedErr.FileName
    line := castedErr.Line
    statusCode := castedErr.StatusCode

    unwrapErr := castedErr.Unwrap()
    errMessage  := unwrapErr.Error()
}

// Create custom http status error
err = httperror.NewError(301, errors.New("error message")).Wrap()

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
