package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RequestValidator[T any] struct {
	validateBody T
}

func (v RequestValidator[T]) GetBodyValidator(c *gin.Context) {
	if err := c.ShouldBindBodyWith(&v.validateBody, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Loading Body failed": err.Error(),
		})
		c.Abort()
		return
	}

	v.validate(c)
}

func (v RequestValidator[T]) GetQueryValidator(c *gin.Context) {
	if err := c.ShouldBindQuery(&v.validateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Loading Query failed": err.Error(),
		})
		c.Abort()
		return
	}

	v.validate(c)
}

func (v RequestValidator[T]) validate(c *gin.Context) {
	validate := validator.New()

	if err := validate.Struct(&v.validateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	c.Next()
}
