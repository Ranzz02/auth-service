package utils

import "net/http"

type ApiError struct {
	Code       uint16 `gorm:"" json:"code"`
	Message    string `gorm:"" json:"error"`
	Details    string `gorm:"" json:"details,omitempty"`
	StatusCode int    `gorm:"-" json:"-"`
}

func (e ApiError) Error() string {
	return e.Message
}

func NewApiError(Code uint16, Error string, Details string, StatusCode int) ApiError {
	return ApiError{
		Code:       Code,
		Message:    Error,
		Details:    Details,
		StatusCode: StatusCode,
	}
}

const (
	// Authentication Errors
	CodeInvalidUsernameOrEmail = 101
	CodeUserIsNotAuthorized    = 102

	// Validation Errors
	CodeUsernameOrEmailInUse = 201
	CodeInvalidInputData     = 202
	CodeInvalidCredentials   = 203

	// Resource Errors
	CodeResourceNotFound = 301
	CodeResourceConflict = 302

	// Server Errors
	CodeInternalServerError = 501
)

var (
	// Authentication errors
	InvalidUsernameOrEmail = NewApiError(
		CodeInvalidUsernameOrEmail,
		"Invalid input. Please check your details and try again.",
		"",
		http.StatusUnauthorized,
	)
	UserIsNotAuthorized = NewApiError(
		CodeUserIsNotAuthorized,
		"User is not authorized to perform this action.",
		"",
		http.StatusForbidden,
	)

	// Validation errors
	UsernameOrEmailInUse = NewApiError(
		CodeUsernameOrEmailInUse,
		"The provided credentials could not be used. Please try again.",
		"",
		http.StatusConflict,
	)
	InvalidInputData = NewApiError(
		CodeInvalidInputData,
		"The input data provided is invalid.",
		"Please review the input fields and try again.",
		http.StatusBadRequest,
	)
	InvalidCredentials = NewApiError(
		CodeInvalidCredentials,
		"Invalid credentials. Please check your username/email and password and try again.",
		"",
		http.StatusBadRequest,
	)

	// Resource errors
	ResourceNotFound = NewApiError(
		CodeResourceNotFound,
		"The requested resource was not found.",
		"",
		http.StatusNotFound,
	)
	ResourceConflict = NewApiError(
		CodeResourceConflict,
		"The resource conflict with an existing one.",
		"",
		http.StatusConflict,
	)

	// Server errors
	InternalServerError = NewApiError(
		CodeInternalServerError,
		"An internal server error occured. Please try again later.",
		"",
		http.StatusInternalServerError,
	)
)
