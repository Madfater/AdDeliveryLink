package middleware

import (
	"net/http"

	"github.com/Madfater/AdDeliveryLink/log"
	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger := log.GetLogger()
				logger.Error("Panic Happened", map[string]interface{}{
					"method":     c.Request.Method,
					"path":       c.Request.URL.Path,
					"user_agent": c.Request.UserAgent(),
					"err":        err,
				})
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
