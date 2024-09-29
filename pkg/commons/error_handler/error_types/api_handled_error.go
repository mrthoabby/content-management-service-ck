package errortypes

import errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"

type ApiHandledError struct {
	error
	message string
}

func (e ApiHandledError) Error() string {
	return e.message
}

func (e ApiHandledError) GetType() uint {
	return errorhandler.ApiHandledError
}

func (e *ApiHandledError) SetTracerLauncherName(name string) {
	e.message = name + ": " + e.message
}

func NewInternalServerError(message string) *ApiHandledError {
	return &ApiHandledError{
		message: message,
	}
}
