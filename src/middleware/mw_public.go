package middleware

import (
	"github.com/Madfater/AdDeliveryLink/src/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PublicMiddleware struct{}

func (PublicMiddleware) PublicQueryValidator(c *gin.Context) {
	var query dto.Query

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Loading Query failed": err.Error()})
		c.Abort()
		return
	}

	validate := validator.New()

	if err := validate.Struct(&query); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.Next()
}
