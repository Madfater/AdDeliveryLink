package routers

import (
	"github.com/Madfater/AdDeliveryLink/controllers"
	_ "github.com/Madfater/AdDeliveryLink/docs"
	"github.com/Madfater/AdDeliveryLink/dto"
	"github.com/Madfater/AdDeliveryLink/middleware"
	"github.com/Madfater/AdDeliveryLink/models"
	"github.com/Madfater/AdDeliveryLink/repositories"
	"github.com/Madfater/AdDeliveryLink/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouterInit(r *gin.Engine) {
	appCtx, _ := models.CreateAppContext("mysql", "redis")
	defer appCtx.Close()

	adsRepo := repositories.NewAdsRepository(appCtx.DB)
	adsService := services.NewAdsService(adsRepo)

	adsController := controllers.NewAdsController(adsService)

	route := r.Group("/v1/api")
	{
		//Admin API
		route.POST("/ad", middleware.RequestValidator[dto.CreateAdsReq]{}.GetBodyValidator, adsController.CreateAdvertisement)

		//Public API
		route.GET("/ad", middleware.RequestValidator[dto.GetAdsResp]{}.GetBodyValidator, adsController.GetAdvertisement)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
