package httperror_test

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/cateiru/go-http-error/httperror"
	"github.com/stretchr/testify/require"
)

func TestCreateError(t *testing.T) {
	p, err := os.Getwd()
	require.NoError(t, err)

	err = httperror.NewError(http.StatusBadRequest, errors.New("test"))

	if castedErr, ok := httperror.CastHTTPError(err); ok {
		require.Equal(t, castedErr.StatusCode, 400)
		require.Equal(t, castedErr.Unwrap().Error(), "test")
		require.Equal(t, castedErr.FileName, fmt.Sprintf("%s/%s", p, "httperror_test.go"))
		require.Equal(t, castedErr.Line, 18)
	}
}

func TestStringError(t *testing.T) {
	p, err := os.Getwd()
	require.NoError(t, err)
	err = httperror.NewStringError(http.StatusInternalServerError, "test")

	if castedErr, ok := httperror.CastHTTPError(err); ok {
		require.Equal(t, castedErr.StatusCode, 500)
		require.Equal(t, castedErr.Unwrap().Error(), "test")
		require.Equal(t, castedErr.FileName, fmt.Sprintf("%s/%s", p, "httperror_test.go"))
		require.Equal(t, castedErr.Line, 31)
	}
}
