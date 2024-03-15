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
		route.POST("/admin", middleware.AdminMiddleware{}.AdminBodyValidator, controllers.AdminController{}.CreateAdvertisement)

		//Public API
		route.GET("/public", middleware.PublicMiddleware{}.PublicQueryValidator, controllers.PublicController{}.PublicAdvertisement)
	}

}
