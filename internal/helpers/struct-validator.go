package helpers

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/hitaloose/go-products-api/internal/exceptions"
)

var validate = validator.New()

func getJSONFieldName(model interface{}, fieldName string) string {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	field, _ := modelType.FieldByName(fieldName)
	return field.Tag.Get("json")
}

func formatError(input interface{}, err error) map[string][]string {
	errorsMap := make(map[string][]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, validationError := range validationErrors {
			fieldName := getJSONFieldName(input, validationError.StructField())
			errorsMap[fieldName] = append(errorsMap[fieldName], validationError.Tag())
		}
	}

	return errorsMap
}

func ValidateStruct(input interface{}) error {
	if err := validate.Struct(input); err != nil {
		return exceptions.NewBadRequestException("Validation failed", formatError(input, err))
	}

	return nil
}
