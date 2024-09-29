package validator

import (
	errortypes "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler/error_types"
	stringutil "github.com/mrthoabby/content-management-service-ck/pkg/util/string_util"
)

func IsNotEmptyString(value, message string) string {
	if stringutil.IsEmptyString(value) {
		panic(errortypes.NewValidationError(message))
	}
	return value
}
