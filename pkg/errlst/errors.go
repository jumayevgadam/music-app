package errlst

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
)

// Package errlst provides custom error handling for HTTP errors.
// It includes various error types and ensures that errors are returned
// with appropriate HTTP status codes, allowing for easier debugging and handling.

// Define common application errors using Go's error package.
// These errors map to HTTP status codes and represent various client and server issues.
var (
	// ErrBadRequest represents an error for a 400 Bad Request response.
	ErrBadRequest = errors.New("bad request")
	// ErrBadQueryParams represents an error for invalid query parameters.
	ErrBadQueryParams = errors.New("bad query params")
	// ErrUnauthorized represents an error for a 401 Unauthorized response.
	ErrUnauthorized = errors.New("unauthorized")
	// ErrForbidden represents an error for a 403 Forbidden response.
	ErrForbidden = errors.New("forbidden")
	// ErrFieldValidation represents an error for invalid field validation.
	ErrFieldValidation = errors.New("field validation error")
	// ErrNotFound represents an error for a 404 Not Found response.
	ErrNotFound = errors.New("not found")
	// ErrNoRecord represents an error when no database record is found.
	ErrNoRecord = errors.New("no records found")
	// ErrMethodNotAllowed represents an error for a 405 Method Not Allowed response.
	ErrMethodNotAllowed = errors.New("method not allowed")
	// ErrNotAcceptable represents an error for a 406 Not Acceptable response.
	ErrNotAcceptable = errors.New("not acceptable")
	// ErrRequestTimeOut represents an error for a 408 Request Timeout response.
	ErrRequestTimeOut = errors.New("request timeout")
	// ErrConflict represents an error for a 409 Conflict response.
	ErrConflict = errors.New("conflict")
	// ErrGone represents an error for a 410 Gone response.
	ErrGone = errors.New("gone")
	// ErrLengthRequired represents an error for missing content length.
	ErrLengthRequired = errors.New("length required")
	// ErrTooManyRequests represents an error for a 429 Too Many Requests response.
	ErrTooManyRequests = errors.New("too many requests")
	// ErrInternalServer represents an error for a 500 Internal Server Error response.
	ErrInternalServer = errors.New("internal server error")
	// ErrServiceUnavailable represents an error for a 503 Service Unavailable response.
	ErrServiceUnavailable = errors.New("service unavailable")
	// ErrTransactionFailed represents an error for transaction failures.
	ErrTransactionFailed = errors.New("failed to perform transaction")
	// ErrNotAllowedImageHeader represents an error for invalid image headers.
	ErrNotAllowedImageHeader = errors.New("not allowed image header")
	// ErrNotRequiredFields represents an error for missing required fields.
	ErrNotRequiredFields = errors.New("missing required fields")
	// ErrRange represents an error for a value being out of range.
	ErrRange = errors.New("value out of range")
	// ErrSyntax represents an error for invalid syntax.
	ErrSyntax = errors.New("invalid syntax")
)

// RestErr defines the interface for custom error handling in the application.
// It allows errors to include an HTTP status code, message, and underlying causes.
type RestErr interface {
	Status() int         // Returns the HTTP status code.
	Error() string       // Returns the error message.
	Causes() interface{} // Returns the underlying cause of the error.
}

// RestError represents a structured error message that implements the RestErr interface.
// It includes the HTTP status code, error message, and any underlying causes of the error.
type RestError struct {
	ErrStatus  int         `json:"err_status,omitempty"`
	ErrMessage string      `json:"err_msg,omitempty"`
	ErrCauses  interface{} `json:"err_cause,omitempty"`
}

// Status returns the HTTP status code associated with the error.
func (e RestError) Status() int {
	return e.ErrStatus
}

// Causes returns the underlying causes of the error.
func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

// Error formats the error message for output, including status code and causes.
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - error_msg: %s - causes: %v", e.ErrStatus, e.ErrMessage, e.ErrCauses)
}

// NewRestError creates a new custom RestError with a given status, message, and causes.
func NewRestError(status int, errMsg string, causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  status,
		ErrMessage: errMsg,
		ErrCauses:  causes,
	}
}

// NewBadRequestError creates a new 400 Bad Request error with the provided cause.
func NewBadRequestError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadRequest.Error(),
		ErrCauses:  causes,
	}
}

// NewNotFoundError creates a new 404 Not Found error with the provided cause.
func NewNotFoundError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusNotFound,
		ErrMessage: ErrNotFound.Error(),
		ErrCauses:  causes,
	}
}

// NewUnAuthorizedError creates a new 401 Unauthorized error with the provided cause.
func NewUnAuthorizedError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusUnauthorized,
		ErrMessage: ErrUnauthorized.Error(),
		ErrCauses:  causes,
	}
}

// NewForbiddenError creates a new 403 Forbidden error with the provided cause.
func NewForbiddenError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusForbidden,
		ErrMessage: ErrForbidden.Error(),
		ErrCauses:  causes,
	}
}

// NewInternalServerError creates a new 500 Internal Server Error with the provided cause.
func NewInternalServerError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: ErrInternalServer.Error(),
		ErrCauses:  causes,
	}
}

// NewRequestTimedOutError creates a new 408 Request Timeout error with the provided cause.
func NewRequestTimedOutError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusRequestTimeout,
		ErrMessage: ErrRequestTimeOut.Error(),
		ErrCauses:  causes,
	}
}

// NewConflictError creates a new 409 Conflict error with the provided cause.
func NewConflictError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusConflict,
		ErrMessage: ErrConflict.Error(),
		ErrCauses:  causes,
	}
}

// NewTooManyRequestError creates a new 429 Too Many Requests error with the provided cause.
func NewTooManyRequestError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusTooManyRequests,
		ErrMessage: ErrTooManyRequests.Error(),
		ErrCauses:  causes,
	}
}

// NewBadQueryParamsError creates a new 400 Bad Request error for invalid query parameters.
func NewBadQueryParamsError(causes interface{}) RestErr {
	return &RestError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadQueryParams.Error(),
		ErrCauses:  causes,
	}
}

// ParseErrors attempts to map a given error to a RestErr.
// It handles various error types such as SQL errors, validation errors, and Go-specific errors.
// If no specific error is matched, it returns a generic Internal Server Error.
func ParseErrors(err error) RestErr {
	switch {
	// Handle Go-specific errors
	case errors.Is(err, pgx.ErrNoRows):
		return NewNotFoundError(ErrNoRecord.Error() + err.Error())
	case errors.Is(err, pgx.ErrTooManyRows):
		return NewConflictError(ErrTooManyRequests.Error() + err.Error())
	case errors.Is(err, context.DeadlineExceeded):
		return NewRequestTimedOutError(err.Error())

	// Handle AWS errors
	case strings.Contains(err.Error(), ErrBucketNotFound.Error()),
		strings.Contains(err.Error(), ErrObjectNotFound.Error()):
		return NewBadRequestError(err.Error())

	// Handle SQLSTATE errors
	case strings.Contains(err.Error(), "SQLSTATE"):
		return ParseSqlErrors(err)

	// Handle strconv.Atoi errors
	case strings.Contains(err.Error(), "invalid syntax"):
		return NewBadRequestError(ErrSyntax.Error())

	// Handle validation errors
	case strings.Contains(err.Error(), ErrFieldValidation.Error()):
		return NewBadRequestError(err.Error())

	// Default case: Return Internal Server Error
	default:
		return NewInternalServerError(err.Error())
	}
}

// parseSQLErrors parses SQL errors and returns corresponding RestErr
func ParseSqlErrors(err error) RestErr {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// CLASS 22
		case "22001": // Numeric value out of range
			return NewBadRequestError("Numeric" + ErrRange.Error() + pgErr.Message)

		// CLASS 23
		case "23505": // Unique violation
			return NewConflictError("Unique constraint violation: " + pgErr.Message)
		case "23503": // Foreign key violation\
			return NewBadRequestError("Foreign key violation: " + pgErr.Message)
		case "23502": // Not-null violation
			return NewBadRequestError("Not-null constraint violation: " + pgErr.Message)
		case "23514": // CHECK violation, a product price must be greater than 0
			return NewBadRequestError("Check violation: " + pgErr.Message)

		// CLASS 40
		case "40001": // Serialization failure
			return NewConflictError("Serialization error: " + pgErr.Message)

		// CLASS 42
		case "42601": // Syntax error
			return NewBadRequestError("Syntax error in SQL statement: " + pgErr.Message)
		}
	}

	if strings.Contains(err.Error(), "scany") {
		return NewBadRequestError(err.Error())
	}

	if strings.Contains(err.Error(), "no corresponding field found") {
		return NewBadRequestError(err.Error())
	}

	return NewBadRequestError(err.Error())
}
