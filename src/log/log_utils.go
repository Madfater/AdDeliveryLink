package log

import (
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse[T any](c *gin.Context, data data.IResponse[T]) {
	logger := GetLogger()
	requestId, _ := c.Get("requestId")
	logger.Info("Request Finished", map[string]interface{}{
		"request_id": requestId,
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"user_agent": c.Request.UserAgent(),
		"response":   data,
	})
	c.JSON(http.StatusOK, data)
}

func ErrorResponse(c *gin.Context, httpStatus int, message string, err error) {
	logger := GetLogger()
	requestId, _ := c.Get("requestId")
	logger.Error("Request Failed", map[string]interface{}{
		"request_id": requestId,
		"method":     c.Request.Method,
		"path":       c.Request.URL.Path,
		"user_agent": c.Request.UserAgent(),
		"err":        err.Error(),
	})
	c.JSON(httpStatus, gin.H{
		"status":  "error",
		"message": message,
		"err":     err.Error(),
	})
}

func HandleError(err error, msg string) {
	logger := GetLogger()
	logger.Error(msg, map[string]interface{}{
		"err": err.Error(),
	})
}
