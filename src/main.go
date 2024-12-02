package main

import (
	"github.com/Madfater/AdDeliveryLink/controllers"
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	_ "github.com/Madfater/AdDeliveryLink/docs"
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

// feature
// TODO: 資料匯入匯出
// TODO: 定期過期廣告
// TODO: 定期清除過期廣告
// TODO: 重新實作 repository test(因為測虛擬的約等於沒用 orm)

// redis
// TODO: 實作 緩存預熱
// TODO: 評估是否有緩存異常

// docs
// TODO: 重新設定 swagger
// TODO: 修改 readme
// TODO: 流程圖等圖

// log
// TODO: 實作 log 機制
// TODO: logstash + ES
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

	validator := utils.ValidationRegister{}
	validator.EnumRegister()

	route := r.Group("/v1/api")
	{
		route.POST("/ad", middleware.NewValidator(data.CreateAdsReq{}).GetBodyValidator, adsController.CreateAdvertisement)

		route.GET("/ad", middleware.NewValidator(data.GetAdsReq{}).GetQueryValidator, adsController.GetAdvertisement)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = r.Run(":8080")
	utils.HandleError(err, "Fail to run server")
}
