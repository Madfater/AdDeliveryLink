package main

import (
	"github.com/Madfater/AdDeliveryLink/controllers"
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"github.com/Madfater/AdDeliveryLink/middleware"
	"github.com/Madfater/AdDeliveryLink/models"
	"github.com/Madfater/AdDeliveryLink/repositories"
	"github.com/Madfater/AdDeliveryLink/services"
	"github.com/Madfater/AdDeliveryLink/utils"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
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
	r.Use(gin.RecoveryWithWriter(io.MultiWriter(f)))
	r.Use(gin.LoggerWithWriter(io.MultiWriter(f)))

	appCtx, err := models.CreateAppContext("mysql", "redis")
	utils.HandleError(err, "Fail to create application context")
	defer appCtx.Close()

	adsRepo := repositories.NewAdsRepository(appCtx.DB)
	adsService := services.NewAdsService(adsRepo)

	adsController := controllers.NewAdsController(adsService)

	route := r.Group("/v1/api")
	{
		//Admin API
		route.POST("/ad", middleware.NewValidator(data.CreateAdsReq{}).GetBodyValidator, adsController.CreateAdvertisement)

		//Public API
		route.GET("/ad", middleware.NewValidator(data.GetAdsReq{}).GetQueryValidator, adsController.GetAdvertisement)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = r.Run(":8080")
	utils.HandleError(err, "Fail to run server")
}
