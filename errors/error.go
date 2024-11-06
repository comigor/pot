package errors

import (
	"errors"
	"net/http"
)

var (
	ErrGeneralBadRequest                    = errors.New("general error, Bad Request")
	ErrGeneralUnauthorized                  = errors.New("general error, Unauthorized")
	ErrGeneralPaymentRequired               = errors.New("general error, Payment Required")
	ErrGeneralForbidden                     = errors.New("general error, Forbidden")
	ErrGeneralNotFound                      = errors.New("general error, Not Found")
	ErrGeneralMethodNotAllowed              = errors.New("general error, Method Not Allowed")
	ErrGeneralNotAcceptable                 = errors.New("general error, Not Acceptable")
	ErrGeneralProxyAuthRequired             = errors.New("general error, Proxy Authentication Required")
	ErrGeneralRequestTimeout                = errors.New("general error, Request Timeout")
	ErrGeneralConflict                      = errors.New("general error, Conflict")
	ErrGeneralGone                          = errors.New("general error, Gone")
	ErrGeneralLengthRequired                = errors.New("general error, Length Required")
	ErrGeneralPreconditionFailed            = errors.New("general error, Precondition Failed")
	ErrGeneralRequestEntityTooLarge         = errors.New("general error, Request Entity Too Large")
	ErrGeneralRequestURITooLong             = errors.New("general error, Request URI Too Long")
	ErrGeneralUnsupportedMediaType          = errors.New("general error, Unsupported Media Type")
	ErrGeneralRequestedRangeNotSatisfiable  = errors.New("general error, Requested Range Not Satisfiable")
	ErrGeneralExpectationFailed             = errors.New("general error, Expectation Failed")
	ErrGeneralTeapot                        = errors.New("general error, I'm a teapot")
	ErrGeneralMisdirectedRequest            = errors.New("general error, Misdirected Request")
	ErrGeneralUnprocessableEntity           = errors.New("general error, Unprocessable Entity")
	ErrGeneralLocked                        = errors.New("general error, Locked")
	ErrGeneralFailedDependency              = errors.New("general error, Failed Dependency")
	ErrGeneralTooEarly                      = errors.New("general error, Too Early")
	ErrGeneralUpgradeRequired               = errors.New("general error, Upgrade Required")
	ErrGeneralPreconditionRequired          = errors.New("general error, Precondition Required")
	ErrGeneralTooManyRequests               = errors.New("general error, Too Many Requests")
	ErrGeneralRequestHeaderFieldsTooLarge   = errors.New("general error, Request Header Fields Too Large")
	ErrGeneralUnavailableForLegalReasons    = errors.New("general error, Unavailable For Legal Reasons")
	ErrGeneralInternalServerError           = errors.New("general error, Internal Server Error")
	ErrGeneralNotImplemented                = errors.New("general error, Not Implemented")
	ErrGeneralBadGateway                    = errors.New("general error, Bad Gateway")
	ErrGeneralServiceUnavailable            = errors.New("general error, Service Unavailable")
	ErrGeneralGatewayTimeout                = errors.New("general error, Gateway Timeout")
	ErrGeneralHTTPVersionNotSupported       = errors.New("general error, HTTP Version Not Supported")
	ErrGeneralVariantAlsoNegotiates         = errors.New("general error, Variant Also Negotiates")
	ErrGeneralInsufficientStorage           = errors.New("general error, Insufficient Storage")
	ErrGeneralLoopDetected                  = errors.New("general error, Loop Detected")
	ErrGeneralNotExtended                   = errors.New("general error, Not Extended")
	ErrGeneralNetworkAuthenticationRequired = errors.New("general error, Network Authentication Required")
)

// Error map to associate status codes with error variables
var httpErrorMap = map[int]error{
	http.StatusBadRequest:                    ErrGeneralBadRequest,
	http.StatusUnauthorized:                  ErrGeneralUnauthorized,
	http.StatusPaymentRequired:               ErrGeneralPaymentRequired,
	http.StatusForbidden:                     ErrGeneralForbidden,
	http.StatusNotFound:                      ErrGeneralNotFound,
	http.StatusMethodNotAllowed:              ErrGeneralMethodNotAllowed,
	http.StatusNotAcceptable:                 ErrGeneralNotAcceptable,
	http.StatusProxyAuthRequired:             ErrGeneralProxyAuthRequired,
	http.StatusRequestTimeout:                ErrGeneralRequestTimeout,
	http.StatusConflict:                      ErrGeneralConflict,
	http.StatusGone:                          ErrGeneralGone,
	http.StatusLengthRequired:                ErrGeneralLengthRequired,
	http.StatusPreconditionFailed:            ErrGeneralPreconditionFailed,
	http.StatusRequestEntityTooLarge:         ErrGeneralRequestEntityTooLarge,
	http.StatusRequestURITooLong:             ErrGeneralRequestURITooLong,
	http.StatusUnsupportedMediaType:          ErrGeneralUnsupportedMediaType,
	http.StatusRequestedRangeNotSatisfiable:  ErrGeneralRequestedRangeNotSatisfiable,
	http.StatusExpectationFailed:             ErrGeneralExpectationFailed,
	http.StatusTeapot:                        ErrGeneralTeapot,
	http.StatusMisdirectedRequest:            ErrGeneralMisdirectedRequest,
	http.StatusUnprocessableEntity:           ErrGeneralUnprocessableEntity,
	http.StatusLocked:                        ErrGeneralLocked,
	http.StatusFailedDependency:              ErrGeneralFailedDependency,
	http.StatusTooEarly:                      ErrGeneralTooEarly,
	http.StatusUpgradeRequired:               ErrGeneralUpgradeRequired,
	http.StatusPreconditionRequired:          ErrGeneralPreconditionRequired,
	http.StatusTooManyRequests:               ErrGeneralTooManyRequests,
	http.StatusRequestHeaderFieldsTooLarge:   ErrGeneralRequestHeaderFieldsTooLarge,
	http.StatusUnavailableForLegalReasons:    ErrGeneralUnavailableForLegalReasons,
	http.StatusInternalServerError:           ErrGeneralInternalServerError,
	http.StatusNotImplemented:                ErrGeneralNotImplemented,
	http.StatusBadGateway:                    ErrGeneralBadGateway,
	http.StatusServiceUnavailable:            ErrGeneralServiceUnavailable,
	http.StatusGatewayTimeout:                ErrGeneralGatewayTimeout,
	http.StatusHTTPVersionNotSupported:       ErrGeneralHTTPVersionNotSupported,
	http.StatusVariantAlsoNegotiates:         ErrGeneralVariantAlsoNegotiates,
	http.StatusInsufficientStorage:           ErrGeneralInsufficientStorage,
	http.StatusLoopDetected:                  ErrGeneralLoopDetected,
	http.StatusNotExtended:                   ErrGeneralNotExtended,
	http.StatusNetworkAuthenticationRequired: ErrGeneralNetworkAuthenticationRequired,
}
