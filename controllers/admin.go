package controllers

import (
	"dcardAssignment/models"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AdminController struct{}

type Body struct {
	Title      string     `json:"title" validate:"required"`
	StartAt    time.Time  `json:"startAt" validate:"required"`
	EndAt      time.Time  `json:"endAt" validate:"required"`
	Conditions Conditions `json:"conditions" validate:"required"`
}

type Conditions struct {
	AgeStart int      `json:"ageStart" validate:"min=1,max=100"`
	AgeEnd   int      `json:"ageEnd" validate:"min=1,max=100"`
	Country  []string `json:"country"`
	Platform []string `json:"platform"`
	Gender   *string  `json:"gender" `
}

func (AdminController) CreateAdvertisement(c *gin.Context) {

	var body Body

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Loadin Body failed": err.Error()})
		return
	}

	if body.Conditions.AgeStart == 0 {
		body.Conditions.AgeStart = 1
	}

	if body.Conditions.AgeEnd == 0 {
		body.Conditions.AgeEnd = 100
	}

	validate := validator.New()

	err := validate.Struct(body)
	if err != nil {
		var errMsg = "=== error msg ====\n"
		errMsg += fmt.Sprintf("%s\n", err)

		if _, ok := err.(*validator.InvalidValidationError); ok {
			errMsg += fmt.Sprintf("%s\n", err)
			return
		}

		errMsg += "=========== error field info ====================\n"
		for _, err := range err.(validator.ValidationErrors) {
			errMsg += fmt.Sprintf("Namespace: %s", err.Namespace())
			errMsg += fmt.Sprintf("Fild: %s", err.Field())
			errMsg += fmt.Sprintf("StructNamespace: %s", err.StructNamespace())
			errMsg += fmt.Sprintf("StructField: %s", err.StructField())
			errMsg += fmt.Sprintf("Tag: %s", err.Tag())
			errMsg += fmt.Sprintf("ActualTag: %s", err.ActualTag())
			errMsg += fmt.Sprintf("Kind: %s", err.Kind())
			errMsg += fmt.Sprintf("Type: %s", err.Type())
			errMsg += fmt.Sprintf("Value: %s", err.Value())
			errMsg += fmt.Sprintf("Param: %s", err.Param())
		}
		c.JSON(http.StatusBadRequest, gin.H{"Error Message": errMsg})
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
