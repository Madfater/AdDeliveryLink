package utils

import (
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"github.com/Madfater/AdDeliveryLink/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse[T any](c *gin.Context, data data.GenericResponse[T]) {
	logger := log.GetLogger()
	logger.Info("Request Finished", map[string]interface{}{
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"user_agent": c.Request.UserAgent(),
		"response":   data,
	})
	c.JSON(http.StatusOK, data)
}

func ErrorResponse(c *gin.Context, httpStatus int, message string, err error) {
	logger := log.GetLogger()
	logger.Error("Request Failed", map[string]interface{}{
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"user_agent": c.Request.UserAgent(),
		"err":        err.Error(),
	})
	c.JSON(httpStatus, gin.H{
		"status":  "error",
		"message": message,
	})
}
