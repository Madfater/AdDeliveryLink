package middleware

import (
	"bytes"
	"github.com/Madfater/AdDeliveryLink/log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := log.GetLogger()

		method := c.Request.Method
		requestId := uuid.New().String()
		c.Set("requestId", requestId)

		info := map[string]interface{}{
			"request_id": requestId,
			"method":     method,
			"path":       c.Request.URL.Path,
			"user_agent": c.Request.UserAgent(),
		}

		switch method {
		case http.MethodGet:
			query := c.Request.URL.RawQuery

			info["query"] = query
		default:
			var requestBody string

			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				log.HandleError(err, "Failed to read request body")
				return
			}

			requestBody = string(bodyBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			info["request_Body"] = requestBody
		}

		logger.Info("Request start", info)

		c.Next()
	}
}
