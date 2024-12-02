package main

import (
	"github.com/Madfater/AdDeliveryLink/controllers"
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	_ "github.com/Madfater/AdDeliveryLink/docs"
	"github.com/Madfater/AdDeliveryLink/enum"
	"github.com/Madfater/AdDeliveryLink/middleware"
	"github.com/Madfater/AdDeliveryLink/models"
	"github.com/Madfater/AdDeliveryLink/repositories"
	"github.com/Madfater/AdDeliveryLink/services"
	"github.com/Madfater/AdDeliveryLink/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"

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

// bug
// TODO: 重新整理儲存的邏輯

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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("enum", ValidateEnum)
		if err != nil {
			utils.HandleError(err, "Fail to register validation")
		}
	}

	route := r.Group("/v1/api")
	{
		route.POST("/ad", middleware.NewValidator(data.CreateAdsReq{}).GetBodyValidator, adsController.CreateAdvertisement)

		route.GET("/ad", middleware.NewValidator(data.GetAdsReq{}).GetQueryValidator, adsController.GetAdvertisement)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = r.Run(":8080")
	utils.HandleError(err, "Fail to run server")
}

func ValidateEnum(fl validator.FieldLevel) bool {

	fieldValue := fl.Field()
	fieldType := fl.Field().Type()

	if fieldType.Kind() == reflect.Slice {
		for i := 0; i < fieldValue.Len(); i++ {
			elem := fieldValue.Index(i).Interface()
			if !validateEnumInterface(elem) {
				return false
			}
		}
		return true
	}

	return validateEnumInterface(fieldValue.Interface())
}

func validateEnumInterface(value interface{}) bool {
	enumInterfaceType := reflect.TypeOf((*enum.EnumInterface)(nil)).Elem()
	valueType := reflect.TypeOf(value)

	if valueType.Implements(enumInterfaceType) {
		if enum, ok := value.(enum.EnumInterface); ok {
			return enum.IsValidEnum()
		}
	}

	return false
}
