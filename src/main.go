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
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Advertisement API
// @version 1.0
// @description a simple API for advertisement

// @BasePath /v1/api
func main() {
	gin.DisableConsoleColor()

	r := gin.New()

	//設定 middleware
	r.Use(middleware.CustomRecovery())
	r.Use(middleware.RequestLogger())
	middleware.RegisterCustomValidation()

	appCtx, err := models.CreateAppContext("mysql", "redis")
	if err != nil {
		log.HandleError(err, "Fail to create application context")
	}
	defer func(appCtx *models.AppContext) {
		err := appCtx.Close()
		log.HandleError(err, "Fail to close application context")
	}(appCtx)

	//依賴注入
	adsRepo := repositories.NewAdsRepository(appCtx.DB)
	adsService := services.NewAdsService(adsRepo)
	adsController := controllers.NewAdsController(adsService)

	logger := log.GetLogger()
	logger.Info("Application start", map[string]interface{}{})

	route := r.Group("/v1/api")
	{
		route.POST("/ad", middleware.NewValidator(data.CreateAdsReq{}).GetValidator, adsController.CreateAdvertisement)

		route.GET("/ad", middleware.NewValidator(data.GetAdsReq{}).GetValidator, adsController.GetAdvertisement)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = r.Run(":8080")
	if err != nil {
		log.HandleError(err, "Fail to run server")
	}
}
