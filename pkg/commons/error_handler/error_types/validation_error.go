package errortypes

import errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"

type ValidationError struct {
	error
	message string
}

func (e ValidationError) Error() string {
	return e.message
}

func (e ValidationError) GetType() uint {
	return errorhandler.ValidationErrorType
}

func (e *ValidationError) SetTracerLauncherName(name string) {
	e.message = name + ": " + e.message
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		message: message,
	}
}
