package main

import (
	"github.com/Madfater/AdDeliveryLink/controllers"
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	_ "github.com/Madfater/AdDeliveryLink/docs"
	"github.com/Madfater/AdDeliveryLink/log"
	"github.com/Madfater/AdDeliveryLink/middleware"
	"github.com/Madfater/AdDeliveryLink/models"
	"github.com/Madfater/AdDeliveryLink/repositories"
	"github.com/Madfater/AdDeliveryLink/services"
	"github.com/Madfater/AdDeliveryLink/utils"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @title Advertisement API
// @version 1.0
// @description a simple API for advertisement

// @BasePath /v1/api
func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")

	r := gin.New()
	r.Use(CustomRecovery())
	r.Use(gin.LoggerWithWriter(io.MultiWriter(f)))

	logger := log.GetLogger()

	appCtx, err := models.CreateAppContext("mysql", "redis")
	utils.HandleError(err, "Fail to create application context")
	defer appCtx.Close()

	adsRepo := repositories.NewAdsRepository(appCtx.DB)
	adsService := services.NewAdsService(adsRepo)

	adsController := controllers.NewAdsController(adsService)

	validator := utils.NewValidationRegistrant()
	validator.RegisterEnum()

	logger.Info("Application start", map[string]interface{}{})

	route := r.Group("/v1/api")
	{
		route.POST("/ad", middleware.NewValidator(data.CreateAdsReq{}).GetBodyValidator, adsController.CreateAdvertisement)

		route.GET("/ad", middleware.NewValidator(data.GetAdsReq{}).GetQueryValidator, adsController.GetAdvertisement)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = r.Run(":8080")
	utils.HandleError(err, "Fail to run server")
}

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
