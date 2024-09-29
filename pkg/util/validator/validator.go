package validator

import (
	"strconv"

	errortypes "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler/error_types"
	stringutil "github.com/mrthoabby/content-management-service-ck/pkg/util/string_util"
)

func IsNotEmptyString(value, message string) string {
	if stringutil.IsEmptyString(value) {
		panic(errortypes.NewValidationError(message + " cannot be empty"))
	}
	return value
}

func IsAValidNumber(value, message string) int {
	possibleNumber := IsNotEmptyString(value, message)
	number, errorGettingNumber := strconv.Atoi(possibleNumber)
	if errorGettingNumber != nil {
		panic(errortypes.NewValidationError(message + " must be a valid number"))
	}
	return number
}

func IsAValidBoolean(value, message string) bool {
	possibleBoolean := IsNotEmptyString(value, message)
	boolean, errorGettingBoolean := strconv.ParseBool(possibleBoolean)
	if errorGettingBoolean != nil {
		panic(errortypes.NewValidationError(message + " must be a valid boolean"))
	}
	return boolean
}
