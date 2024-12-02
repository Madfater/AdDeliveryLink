package middleware

import (
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"github.com/Madfater/AdDeliveryLink/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type RequestValidator[T data.ReqInterface] struct {
	requestType T
}

func NewValidator[T data.ReqInterface](requestType T) *RequestValidator[T] {
	return &RequestValidator[T]{requestType: requestType}
}

func (v *RequestValidator[T]) GetBodyValidator(c *gin.Context) {
	if err := c.ShouldBindBodyWith(&v.requestType, binding.JSON); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		c.Abort()
		return
	}

	v.validate(c)
}

func (v *RequestValidator[T]) GetQueryValidator(c *gin.Context) {
	if err := c.ShouldBindQuery(&v.requestType); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters")
		c.Abort()
		return
	}

	v.validate(c)
}

func (v *RequestValidator[T]) validate(c *gin.Context) {
	if isValid, err := v.requestType.Validate(); err != nil && !isValid {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	c.Next()
}
