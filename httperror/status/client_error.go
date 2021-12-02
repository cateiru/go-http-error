package status

import (
	"net/http"

	"github.com/cateiru/go-http-error/httperror"
)

// 400 Bad Request
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400
func NewBadRequestError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusBadRequest, err)
}

// 401 Unauthorized
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401
func NewUnauthorizedError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusUnauthorized, err)
}

// 402 Payment Required
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/402
func NewPaymentRequiredError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusPaymentRequired, err)
}

// 403 Forbidden
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403
func NewForbiddenError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusForbidden, err)
}

// 404 Not Found
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404
func NewNotFoundError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusNotFound, err)
}

// 405 Method Not Allowed
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
func NewMethodNotAllowedError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusMethodNotAllowed, err)
}

// 406 Not Acceptable
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406
func NewNotAcceptableError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusNotAcceptable, err)
}

// 407 Proxy Authentication Required
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/407
func NewProxyAuthenticationRequiredError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusProxyAuthRequired, err)
}

// 408 Request Timeout
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408
func NewRequestTimeoutError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusRequestTimeout, err)
}

// 409 Conflict
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409
func NewConflictError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusConflict, err)
}

// 410 Gone
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/410
func NewGoneError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusGone, err)
}

// 411 Length Required
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/411
func NewLengthRequiredError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusLengthRequired, err)
}

// 412 Precondition Failed
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/412
func NewPreconditionFailedError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusPreconditionFailed, err)
}

// 413 Payload Too Large
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413
func NewPayloadTooLargeError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusRequestEntityTooLarge, err)
}

// 414 URI Too Long
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/414
func NewURITooLongError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusRequestURITooLong, err)
}

// 415 Unsupported Media Type
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415
func NewUnsupportedMediaTypeError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusUnsupportedMediaType, err)
}

// 416 Range Not Satisfiable
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/416
func NewRangeNotSatisfiableError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusRequestedRangeNotSatisfiable, err)
}

// 417 Expectation Failed
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/417
func NewExpectationFailedError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusExpectationFailed, err)
}

// 418 I'm a teapot
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/418
func NewTeapotError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusTeapot, err)
}

// 422 Unprocessable Entity
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422
func NewUnprocessableEntityError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusUnprocessableEntity, err)
}

// 425 Too Early
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/425
func NewTooEarlyError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusTooEarly, err)
}

// 426 Upgrade Required
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/426
func NewUpgradeRequiredError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusUpgradeRequired, err)
}

// 428 Precondition Required
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/428
func NewPreconditionRequiredError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusPreconditionRequired, err)
}

// 429 Too Many Requests
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429
func NewTooManyRequestsError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusTooManyRequests, err)
}

// 431 Request Header Fields Too Large
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/431
func NewRequestHeaderFieldsTooLargeError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusRequestHeaderFieldsTooLarge, err)
}

// 451 Unavailable For Legal Reasons
//
// See more: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/451
func NewUnavailableForLegalReasonsError(err error) *httperror.HTTPError {
	return httperror.NewError(http.StatusUnavailableForLegalReasons, err)
}
