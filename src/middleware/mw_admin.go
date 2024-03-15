package middleware

import (
	"dcardAssignment/src/dto"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type AdminMiddleware struct{}

func (AdminMiddleware) AdminBodyValidator(c *gin.Context) {
	var body dto.Body

	if err := c.ShouldBindBodyWith(&body,binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Loading Body failed": err.Error()})
		return
	}

	validate := validator.New()

	if err := validate.Struct(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.Next()
}
