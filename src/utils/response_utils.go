package utils

import (
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse[T any](c *gin.Context, data data.GenericResponse[T]) {
	c.JSON(http.StatusOK, data)
}

func ErrorResponse(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, gin.H{
		"status":  "error",
		"message": message,
	})
}
