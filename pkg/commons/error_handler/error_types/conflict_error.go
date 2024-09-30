package errortypes

import errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"

func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		message: message,
	}
}

type ConflictError struct {
	error
	message string
}

func (e ConflictError) Error() string {
	return e.message
}

func (e ConflictError) GetType() uint {
	return errorhandler.ConflictErrorType
}

func (e *ConflictError) SetTracerLauncherName(name string) {
	e.message = name + ": " + e.message
}
