package controllers

import (
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"github.com/Madfater/AdDeliveryLink/log"
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
// @Success 200 {object} data.IResponse[entity.Advertisement]
// @Failure 400
// @Router /ad [post]
func (ctrl *AdsController) CreateAdvertisement(c *gin.Context) {
	var body data.CreateAdsReq

	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		log.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	logger := log.GetLogger()
	logger.Info("Request processed", map[string]interface{}{
		"method":       c.Request.Method,
		"path":         c.Request.URL.Path,
		"user_agent":   c.Request.UserAgent(),
		"request_body": body,
	})

	response, err := ctrl.service.CreateAdvertisement(body)
	if err != nil {
		log.ErrorResponse(c, http.StatusInternalServerError, "Failed to create advertisement", err)
		return
	}

	log.SuccessResponse(c, response)
}

// GetAdvertisement @Summary Gets a list of advertisements
// @Description Gets a list of advertisements that match the specified conditions.
// @Param query query data.GetAdsReq true "Advertisement query parameters"
// @Tags Advertisement
// @Success 200 {object} data.IResponse[data.GetAdsResp]
// @Failure 400
// @Router /ad [get]
func (ctrl *AdsController) GetAdvertisement(c *gin.Context) {
	var query data.GetAdsReq

	if err := c.ShouldBindQuery(&query); err != nil {
		log.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameters", err)
		return
	}

	logger := log.GetLogger()
	logger.Info("Request processed", map[string]interface{}{
		"method":        c.Request.Method,
		"path":          c.Request.URL.Path,
		"user_agent":    c.Request.UserAgent(),
		"request_query": query,
	})

	result, err := ctrl.service.GetAdvertisements(query)
	if err != nil {
		log.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch advertisements", err)
		return
	}

	log.SuccessResponse(c, result)
}
