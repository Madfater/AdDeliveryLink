package routers

import (
	"dcardAssignment/src/controllers"
	"dcardAssignment/src/middleware"
	_ "dcardAssignment/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouterInit(r *gin.Engine) {
	route := r.Group("/v1/api")
	{
		//Admin API
		route.POST("/ad", middleware.AdminMiddleware{}.AdminBodyValidator, controllers.AdminController{}.CreateAdvertisement)

		//Public API
		route.GET("/ad", middleware.PublicMiddleware{}.PublicQueryValidator, controllers.PublicController{}.PublicAdvertisement)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
