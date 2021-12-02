package status

import (
	"net/http"

	"github.com/cateiru/go-http-error/httperror"
)

// 500 Internal Server Error
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500
func NewInternalServerErrorError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusInternalServerError, err)
}

// 501 Not Implemented
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/501
func NewNotImplementedError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusNotImplemented, err)
}

// 502 Bad Gateway
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502
func NewBadGatewayError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusBadGateway, err)
}

// 503 Service Unavailable
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503
func NewServiceUnavailableError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusServiceUnavailable, err)
}

// 504 Gateway Timeout
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504
func NewGatewayTimeoutError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusGatewayTimeout, err)
}

// 505 HTTP Version Not Supported
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/505
func NewHTTPVersionNotSupportedError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusHTTPVersionNotSupported, err)
}

// 506 Variant Also Negotiates
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/506
func NewVariantAlsoNegotiatesError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusVariantAlsoNegotiates, err)
}

// 507 Insufficient Storage
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/507
func NewInsufficientStorageError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusInsufficientStorage, err)
}

// 508 Loop Detected
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/508
func NewLoopDetectedError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusLoopDetected, err)
}

// 510 Not Extended
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/510
func NewNotExtendedError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusNotExtended, err)
}

// 511 Network Authentication Required
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/511
func NewNetworkAuthenticationRequiredError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusNetworkAuthenticationRequired, err)
}
