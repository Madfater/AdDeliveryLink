package controlers

import (
	"dcardAssignment/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PublicControler struct{}

type Item struct {
	Title string    `json:"Title"`
	EndAt time.Time `json:"EndAt"`
}

func (PublicControler) PublicAdvertisement(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	age, _ := strconv.Atoi(c.Query("age"))
	gender := c.Query("gender")
	country := c.Query("country")
	platform := c.Query("platform")

	var selectResult []Item

	models.DB.
		Model(&models.Advertisement{}).
		Select("advertisement.Title", "advertisement_country.CountryCode").
		Joins("left join advertisement_country on advertisement.ID = advertisement_country.advertisementID").
		Joins("left join advertisement_platform on advertisement.ID = advertisement_platform.advertisementID").
		Where(models.DB.Where("advertisement_country.CountryCode = ?", country).Or("advertisement_country.CountryCode is NULL")).
		Where(models.DB.Where("advertisement_platform.PlatformName = ?", platform).Or("advertisement_platform.PlatformName is NULL")).
		Where(models.DB.Where("advertisement.Gender = ?", gender).Or("advertisement.Gender is NULL")).
		Where("advertisement.AgeStart <= ? AND advertisement.AgeEnd >= ?", age, age).
		Offset(offset).
		Limit(limit).
		Order("endAt asc").
		Find(&selectResult)

	c.JSON(http.StatusOK, gin.H{"items": selectResult})
}
