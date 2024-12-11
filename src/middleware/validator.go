package middleware

import (
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"github.com/Madfater/AdDeliveryLink/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type RequestValidator[T data.IRequest] struct {
	requestType T
}

func NewValidator[T data.IRequest](requestType T) *RequestValidator[T] {
	return &RequestValidator[T]{requestType: requestType}
}

func (v *RequestValidator[T]) GetValidator(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		if err := c.ShouldBindQuery(&v.requestType); err != nil {
			log.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err)
			c.Abort()
			return
		}
	default:
		if err := c.ShouldBindBodyWith(&v.requestType, binding.JSON); err != nil {
			log.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
			c.Abort()
			return
		}
	}

	if isValid, err := v.requestType.Validate(); err != nil && !isValid {
		log.ErrorResponse(c, http.StatusBadRequest, "Invalid content", err)
		c.Abort()
		return
	}

	c.Next()
}
