package httperror_test

import (
	"errors"
	"testing"

	"github.com/cateiru/go-http-error/httperror"
	"github.com/stretchr/testify/require"
)

func TestCreateError(t *testing.T) {
	err := httperror.NewInternalServerErrorError(errors.New("test")).Wrap()
	require.Equal(t, err.Error(), "500:: test")

	err = httperror.NewTeapotError(errors.New("test")).Caller("test.go", 10).Wrap()
	require.Equal(t, err.Error(), "\"test.go\", line 10, 418:: test")
}

func TestCastError(t *testing.T) {
	err := httperror.NewInternalServerErrorError(errors.New("test")).Caller("test.go", 10).Wrap()

	if castedErr, ok := httperror.CastHTTPError(err); ok {
		require.Equal(t, castedErr.StatusCode, 500)
		require.Equal(t, castedErr.Unwrap().Error(), "test")
		require.Equal(t, castedErr.FileName, "test.go")
		require.Equal(t, castedErr.Line, 10)
	}
}
