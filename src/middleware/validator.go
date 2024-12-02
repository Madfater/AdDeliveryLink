package middleware

import (
	"github.com/Madfater/AdDeliveryLink/controllers/data"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"Loading Body failed": err.Error(),
		})
		c.Abort()
		return
	}

	v.validate(c)
}

func (v *RequestValidator[T]) GetQueryValidator(c *gin.Context) {
	if err := c.ShouldBindQuery(&v.requestType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Loading Query failed": err.Error(),
		})
		c.Abort()
		return
	}

	v.validate(c)
}

func (v *RequestValidator[T]) validate(c *gin.Context) {
	if isValid, err := v.requestType.Validate(); err != nil && !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"validate error": err.Error(),
		})
		c.Abort()
		return
	}

	c.Next()
}
