package routers

import (
	"dcardAssignment/src/controllers"
	"dcardAssignment/src/middleware"

	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	route := r.Group("/v1/api")
	{
		//Admin API
		route.POST("/ad", middleware.AdminMiddleware{}.AdminBodyValidator, controllers.AdminController{}.CreateAdvertisement)

		//Public API
		route.GET("/ad", middleware.PublicMiddleware{}.PublicQueryValidator, controllers.PublicController{}.PublicAdvertisement)
	}

}
