package middleware

import (
	"github.com/Madfater/AdDeliveryLink/controllers/data/base"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RequestValidator[T base.ReqInterface] struct {
	validateBody T
}

func NewValidator[T base.ReqInterface](validateBody T) *RequestValidator[T] {
	return &RequestValidator[T]{validateBody: validateBody}
}

func (v *RequestValidator[T]) GetBodyValidator(c *gin.Context) {
	if err := c.ShouldBindBodyWith(&v.validateBody, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Loading Body failed": err.Error(),
		})
		c.Abort()
		return
	}

	v.validate(c)
}

func (v *RequestValidator[T]) GetQueryValidator(c *gin.Context) {
	if err := c.ShouldBindQuery(&v.validateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Loading Query failed": err.Error(),
		})
		c.Abort()
		return
	}

	v.validate(c)
}

func (v *RequestValidator[T]) validate(c *gin.Context) {
	if isValid, err := v.validateBody.Validate(); err != nil && !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"validate error": err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()

	if err := validate.Struct(&v.validateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"validate error": err.Error(),
		})
		c.Abort()
		return
	}

	c.Next()
}
