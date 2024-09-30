package errorhandler

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

func Handle(errorReceived error, instance any, asideData ...string) {
	if errorReceived != nil {
		logrus.Error(errorReceived.Error() + " [instance: " + getTypeName(instance) + buildAsideData(asideData...) + "]")
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
	var code int
	var errMsg string

	switch common.GetType() {
	case NotFoundErrorType:
		code = http.StatusNotFound
	case DomainErrorType:
		code = http.StatusUnprocessableEntity
	case ValidationErrorType, InvalidFormatErrorType:
		code = http.StatusBadRequest
	case ApiHandledError:
		code = http.StatusInternalServerError
	case ConflictErrorType:
		code = http.StatusConflict
	default:
		code = http.StatusInternalServerError
		errMsg = "common unknown error"
	}

	if errMsg == "" {
		errMsg = common.Error()
	}

	return HandledError{
		Error: errMsg,
		Code:  code,
	}
}

func buildAsideData(asideData ...string) string {
	if len(asideData) > 0 {
		return "(" + strings.Join(asideData, " ") + ")"
	}
	return ""
}

func getTypeName(value any) string {
	v := reflect.ValueOf(value)

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return "Nil pointer"
		} else {
			return v.Elem().Type().Name()
		}
	}
	return v.Type().Name()
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
