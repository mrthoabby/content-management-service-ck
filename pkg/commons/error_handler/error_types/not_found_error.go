package errortypes

import errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"

type NotFoundError struct {
	error
	message string
}

func (e NotFoundError) Error() string {
	return e.message
}

func (e NotFoundError) GetType() uint {
	return errorhandler.NotFoundErrorType
}

func (e *NotFoundError) SetTracerLauncherName(name string) {
	e.message = name + ": " + e.message
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		message: message,
	}
}
