package utils

import (
	"github.com/Madfater/AdDeliveryLink/enum"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type ValidationRegistrant struct{}

func NewValidationRegistrant() ValidationRegistrant {
	return ValidationRegistrant{}
}

func (v ValidationRegistrant) RegisterEnum() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("enum", validateEnum)
		if err != nil {
			HandleError(err, "Fail to register Enum validation")
		}
	}
}

func validateEnum(fl validator.FieldLevel) bool {

	fieldValue := fl.Field()
	fieldType := fl.Field().Type()

	if fieldType.Kind() == reflect.Slice {
		for i := 0; i < fieldValue.Len(); i++ {
			elem := fieldValue.Index(i).Interface()
			if !validateEnumInterface(elem) {
				return false
			}
		}
		return true
	}

	return validateEnumInterface(fieldValue.Interface())
}

func validateEnumInterface(value interface{}) bool {
	enumInterfaceType := reflect.TypeOf((*enum.EnumInterface)(nil)).Elem()
	valueType := reflect.TypeOf(value)

	if valueType.Implements(enumInterfaceType) {
		if enum, ok := value.(enum.EnumInterface); ok {
			return enum.IsValidEnum()
		}
	}

	return false
}
