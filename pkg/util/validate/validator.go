package validate

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	errortypes "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler/error_types"
	stringutil "github.com/mrthoabby/content-management-service-ck/pkg/util/string_util"
	"github.com/sirupsen/logrus"
)

var validate = validator.New()

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

func IsAValidStructure(value any, messages ...string) {
	if errorValidating := validate.Struct(value); errorValidating != nil {
		logrus.Error(fmt.Sprintf("%s %s", errorValidating, strings.Join(messages, " ")))
		for _, fieldError := range errorValidating.(validator.ValidationErrors) {
			handleCustomValidadorMessages(value, fieldError)
		}
	}
}

func handleCustomValidadorMessages(structure any, fieldError validator.FieldError) {
	fieldName := getJSONFieldName(structure, fieldError.Field())
	var message string

	switch fieldError.Tag() {
	case "required":
		message = fieldName + " is required"
	case "email":
		message = fieldName + " must be a valid email"
	case "url":
		message = fieldName + " must be a valid URL"
	case "numeric":
		message = fieldName + " must be a valid number"
	case "boolean":
		message = fieldName + " must be a valid boolean"
	case "len":
		message = fieldName + " must have a length of " + fieldError.Param()
	case "max":
		message = fieldName + " must be less than or equal to " + fieldError.Param()
	case "min":
		message = fieldName + " must be greater than or equal to " + fieldError.Param()
	case "oneof":
		message = fieldName + " must be one of the following values: " + fieldError.Param()
	case "unique":
		message = fieldName + " must be unique"
	default:
		message = fieldName + " is not valid"
	}

	panic(errortypes.NewValidationError(message))
}

func getJSONFieldName(s any, name string) string {
	reflectedType := reflect.TypeOf(s)
	if reflectedType.Kind() == reflect.Ptr {
		reflectedType = reflectedType.Elem()
	}

	for i := 0; i < reflectedType.NumField(); i++ {
		field := reflectedType.Field(i)
		if field.Name == name {
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" {
				parts := strings.Split(jsonTag, ",")
				return parts[0]
			}
			return field.Name
		}
	}
	return name
}
