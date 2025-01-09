package utils

type ApiError struct {
	Code    uint16 `gorm:"" json:"code"`
	Message string `gorm:"" json:"error"`
	Details string `gorm:"" json:"details,omitempty"`
}

func (e ApiError) Error() string {
	return e.Message
}

func NewApiError(Code uint16, Error string, Details string) ApiError {
	return ApiError{
		Code:    Code,
		Message: Error,
		Details: Details,
	}
}

const (
	// Authentication Errors
	CodeInvalidUsernameOrEmail = 101
	CodeUserIsNotAuthorized    = 102

	// Validation Errors
	CodeUsernameOrEmailInUse = 201
	CodeInvalidInputData     = 202

	// Resource Errors
	CodeResourceNotFound = 301
	CodeResourceConflict = 302

	// Server Errors
	CodeInternalServerError = 501
)

var (
	// Authentication Errors
	InvalidUsernameOrEmail = NewApiError(CodeInvalidUsernameOrEmail, "Invalid input. Please check your details and try again.", "")
)
