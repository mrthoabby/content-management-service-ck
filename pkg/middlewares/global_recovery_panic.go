package middlewares

import (
	"encoding/json"
	"net/http"

	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func GlobalRecoveryPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				handledError := errorhandler.GetHanledError(err)
				responseWriter.Header().Set("Content-Type", "application/json")
				responseWriter.WriteHeader(handledError.Code)
				json.NewEncoder(responseWriter).Encode(handledError)
			}
		}()

		next.ServeHTTP(responseWriter, request)
	})
}
