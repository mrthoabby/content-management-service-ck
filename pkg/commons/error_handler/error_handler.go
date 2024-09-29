package errorhandler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Handle(errorReceived error) {
	if errorReceived != nil {
		logrus.Error(errorReceived.Error())
		commonError, isCommonError := errorReceived.(Commons)
		if isCommonError {
			panic(buildHandledCommonError(commonError))
		} else {
			panic(HandledError{
				Error: "unexpected error",
				Code:  http.StatusInternalServerError,
			})
		}
	}
}

func buildHandledCommonError(common Commons) HandledError {
	switch common.GetType() {
	case NotFoundErrorType:
		return HandledError{
			Error: common.Error(),
			Code:  http.StatusNotFound,
		}
	case DomainErrorType:
		return HandledError{
			Error: common.Error(),
			Code:  http.StatusUnprocessableEntity,
		}
	case ValidationErrorType:
		return HandledError{
			Error: common.Error(),
			Code:  http.StatusBadRequest,
		}
	case ApiHandledError:
		return HandledError{
			Error: common.Error(),
			Code:  http.StatusInternalServerError,
		}
	default:
		return HandledError{
			Error: "common unknown error",
			Code:  http.StatusInternalServerError,
		}
	}

}

func GetHanledError(err any) HandledError {
	handled, isHandled := err.(HandledError)
	common, isCommon := err.(Commons)
	if isHandled {
		return handled
	}

	if isCommon {
		return buildHandledCommonError(common)
	}

	logrus.Error(err)
	return HandledError{
		Error: "unexpected unknown error",
		Code:  http.StatusInternalServerError,
	}
}

type HandledError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}
