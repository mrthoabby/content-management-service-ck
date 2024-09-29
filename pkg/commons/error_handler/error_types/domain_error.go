package errortypes

import errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"

type DomainError struct {
	error
	message string
}

func (e DomainError) Error() string {
	return e.message
}

func (e DomainError) GetType() uint {
	return errorhandler.DomainErrorType
}

func (e *DomainError) SetTracerLauncherName(name string) {
	e.message = name + ": " + e.message
}

func NewDomainError(message string) *DomainError {
	return &DomainError{
		message: message,
	}
}
