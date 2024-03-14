package controllers

import (
	"dcardAssignment/src/models"
	"dcardAssignment/src/dto"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (AdminController) CreateAdvertisement(c *gin.Context) {

	var body dto.Body

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Loadin Body failed": err.Error()})
		return
	}

	ad := models.Advertisement{
		Title:    body.Title,
		StartAt:  body.StartAt,
		EndAt:    body.EndAt,
		AgeStart: body.Conditions.AgeStart,
		AgeEnd:   body.Conditions.AgeEnd,
		Gender:   body.Conditions.Gender,
		Country:  convertToCountries(body.Conditions.Country),
		Platform: convertToPlatforms(body.Conditions.Platform),
	}

	res := models.DB.Create(&ad)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Create failed:": res.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "建立成功"})

}

func convertToPlatforms(platformNames []string) []models.Platform {
	var platforms []models.Platform
	for _, name := range platformNames {
		platforms = append(platforms, models.Platform{PlatformName: name})
	}
	return platforms
}

func convertToCountries(platformNames []string) []models.Country {
	var countries []models.Country
	for _, name := range platformNames {
		countries = append(countries, models.Country{CountryCode: name})
	}
	return countries
}
