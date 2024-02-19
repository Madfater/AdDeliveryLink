package controlers

import (
	"dcardAssignment/models"
	
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminControler struct{}

func (AdminControler) CreateAdvertisement(c *gin.Context) {

	var data []models.Advertisement = models.GetAdList()
	c.JSON(http.StatusOK, data)
}
