package controlers

import (
	"dcardAssignment/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type PublicControler struct{}

type Item struct {
	Title string    `json:"Title"`
	EndAt time.Time `json:"EndAt"`
}

type QueryParams struct {
	Offset   *int   `validate:"required,min=0"`
	Limit    *int   `validate:"required,min=1"`
	Age      *int   `validate:"required,min=1,max=100"`
	Gender   string `validate:"required,oneof=M F"`
	Country  string `validate:"required"`
	Platform string `validate:"required"`
}

func (PublicControler) PublicAdvertisement(c *gin.Context) {
	//讀取query
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	age, _ := strconv.Atoi(c.Query("age"))
	gender := c.Query("gender")
	country := c.Query("country")
	platform := c.Query("platform")

	//驗證query
	validate := validator.New()

	queryParams := QueryParams{
		Offset:   &offset,
		Limit:    &limit,
		Age:      &age,
		Gender:   gender,
		Country:  country,
		Platform: platform,
	}

	err := validate.Struct(queryParams)
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

	var selectResult []Item

	models.DB.
		Model(&models.Advertisement{}).
		Select("advertisement.Title", "advertisement_country.country_code").
		Joins("left join advertisement_country on advertisement.ID = advertisement_country.advertisement_id").
		Joins("left join advertisement_platform on advertisement.ID = advertisement_platform.advertisement_id").
		Where(models.DB.Where("advertisement_country.country_code = ?", country).Or("advertisement_country.country_code is NULL")).
		Where(models.DB.Where("advertisement_platform.platform_name = ?", platform).Or("advertisement_platform.platform_name is NULL")).
		Where(models.DB.Where("advertisement.Gender = ?", gender).Or("advertisement.Gender is NULL")).
		Where("advertisement.AgeStart <= ? AND advertisement.AgeEnd >= ?", age, age).
		Offset(offset).
		Limit(limit).
		Order("endAt asc").
		Find(&selectResult)

	c.JSON(http.StatusOK, gin.H{"items": selectResult})
}
