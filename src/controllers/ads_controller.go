package controllers

import (
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"github.com/Madfater/AdDeliveryLink/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type AdsController struct {
	service services.AdsService
}

func NewAdsController(service services.AdsService) *AdsController {
	return &AdsController{service: service}
}

// CreateAdvertisement @Summary Creates a new advertisement
// @Description Creates a new advertisement with the specified title, start and end dates, and conditions.
// @Param body body data.CreateAdsReq true "Advertisement information"
// @Tags Advertisement
// @Success 200
// @Failure 400
// @Router /ad [post]
func (ctrl *AdsController) CreateAdvertisement(c *gin.Context) {
	var body data.CreateAdsReq

	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	err := ctrl.service.CreateAdvertisement(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "成功建立"})
}

// GetAdvertisement @Summary Gets a list of advertisements
// @Description Gets a list of advertisements that match the specified conditions.
// @Param query query data.GetAdsReq true "Advertisement query parameters"
// @Tags Advertisement
// @Success 200 {object} []data.GetAdsResp
// @Failure 400
// @Router /ad [get]
func (ctrl *AdsController) GetAdvertisement(c *gin.Context) {
	var query data.GetAdsReq

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid query parameters", "error": err.Error()})
		return
	}

	results, err := ctrl.service.GetAdvertisements(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch advertisements", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": results})
}
