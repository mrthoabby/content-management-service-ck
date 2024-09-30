package errortypes

import errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"

func NewInvalidFormatError(message string) *InvalidFormatError {
	return &InvalidFormatError{
		message: message,
	}
}

type InvalidFormatError struct {
	error
	message string
}

func (e InvalidFormatError) Error() string {
	return e.message
}

func (e InvalidFormatError) GetType() uint {
	return errorhandler.InvalidFormatErrorType
}

func (e *InvalidFormatError) SetTracerLauncherName(name string) {
	e.message = name + ": " + e.message
}
