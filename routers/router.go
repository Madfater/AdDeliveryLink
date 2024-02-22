package routers

import (
	"dcardAssignment/controlers"

	"github.com/gin-gonic/gin"
)

func RouterInit(r *gin.Engine) {
	route:=r.Group("/v1/api")
	{
		//Admin API
		route.POST("/admin", controlers.AdminControler{}.CreateAdvertisement)

		//Public API
		route.GET("/public", controlers.PublicControler{}.PublicAdvertisement)
	}

}
