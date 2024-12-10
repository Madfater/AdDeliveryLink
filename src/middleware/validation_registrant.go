package middleware

import (
	"github.com/Madfater/AdDeliveryLink/enum"
	"github.com/Madfater/AdDeliveryLink/log"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// RegisterCustomValidation 註冊自訂 Enum 驗證於 /go-playground/validator
func RegisterCustomValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("enum", enumValidation)
		if err != nil {
			log.HandleError(err, "validator register failed")
		}
	}
}

// enumValidation 校驗字段是否為有效的 Enum 值或 Enum Slice
func enumValidation(fl validator.FieldLevel) bool {
	fieldValue := fl.Field()

	// 如果是 Slice，檢查每個元素
	if fieldValue.Kind() == reflect.Slice {
		for i := 0; i < fieldValue.Len(); i++ {
			if !isEnumValid(fieldValue.Index(i).Interface()) {
				return false
			}
		}
		return true
	}

	// 單一值檢查
	return isEnumValid(fieldValue.Interface())
}

// isEnumValid 檢查值是否實現了 EnumInterface 並有效
func isEnumValid(value interface{}) bool {
	if enum, ok := value.(enum.EnumInterface); ok {
		return enum.IsValidEnum()
	}
	return false
}
