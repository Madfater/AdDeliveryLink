package controllers

import (
	"dcardAssignment/src/dto"
	"dcardAssignment/src/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AdminController struct{}

func (AdminController) CreateAdvertisement(c *gin.Context) {

	var body dto.Body

	if err := c.ShouldBindBodyWith(&body,binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Loading Body failed": err.Error()})
		return
	}

	var defaultAgeStart int = 1
	var defaultAgeEnd int = 100

	if body.Conditions.AgeStart == nil {
		body.Conditions.AgeStart = &defaultAgeStart
	}

	if body.Conditions.AgeEnd == nil {
		body.Conditions.AgeEnd = &defaultAgeEnd
	}

	ad := models.Advertisement{
		Title:    body.Title,
		StartAt:  body.StartAt,
		EndAt:    body.EndAt,
		AgeStart: *body.Conditions.AgeStart,
		AgeEnd:   *body.Conditions.AgeEnd,
		Gender:   body.Conditions.Gender,
		Country:  convertToCountries(body.Conditions.Country),
		Platform: convertToPlatforms(body.Conditions.Platform),
	}

	res := models.DB.Debug().Create(&ad)
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
