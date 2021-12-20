package httperror_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/cateiru/go-http-error/httperror"
	"github.com/cateiru/go-http-error/httperror/status"
	"github.com/stretchr/testify/require"
)

func TestCreateError(t *testing.T) {
	p, err := os.Getwd()
	require.NoError(t, err)
	err = status.NewInternalServerErrorError(errors.New("test"))
	require.Equal(t, err.Error(), "500:: test")

	err = status.NewTeapotError(errors.New("test")).Caller()
	require.Equal(t, err.Error(), fmt.Sprintf("\"%s/%s\", line %d, 418:: test", p, "httperror_test.go", 20))
}

func TestCastError(t *testing.T) {
	p, err := os.Getwd()
	require.NoError(t, err)
	err = status.NewInternalServerErrorError(errors.New("test")).Caller().AddCode(10)

	if castedErr, ok := httperror.CastHTTPError(err); ok {
		require.Equal(t, castedErr.StatusCode, 500)
		require.Equal(t, castedErr.Unwrap().Error(), "test")
		require.Equal(t, castedErr.FileName, fmt.Sprintf("%s/%s", p, "httperror_test.go"))
		require.Equal(t, castedErr.Code, 10)
		require.Equal(t, castedErr.Line, 27)
	}
}
