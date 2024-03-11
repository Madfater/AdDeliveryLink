package routers

import (
	"dcardAssignment/controllers"

	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	route:=r.Group("/v1/api")
	{
		//Admin API
		route.POST("/admin", controllers.AdminController{}.CreateAdvertisement)

		//Public API
		route.GET("/public", controllers.PublicController{}.PublicAdvertisement)
	}

}
