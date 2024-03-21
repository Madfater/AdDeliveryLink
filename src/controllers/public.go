package controllers

import (
	"dcardAssignment/src/dto"
	"dcardAssignment/src/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PublicController struct{}

// @Summary Gets a list of advertisements
// @Description Gets a list of advertisements that match the specified conditions.
// @Param query query dto.Query true "Advertisement query parameters"
// @Tags Advertisement
// @Success 200 {object} []dto.Response
// @Failure 400
// @Router /ad [get]
func (PublicController) PublicAdvertisement(c *gin.Context) {

	var query dto.Query

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": err})
	}

	var defaultLimitValue int = 5
	var defaultOffsetValue int = 0

	if query.Limit == nil {
		query.Limit = &defaultLimitValue
	}

	if query.Offset == nil {
		query.Offset = &defaultOffsetValue
	}

	var selectResult []dto.Response

	sqlQuery := models.DB.
		Model(&models.Advertisement{}).
		Select("DISTINCT advertisement.Title", "advertisement.EndAt").
		Joins("left join advertisement_country on advertisement.ID = advertisement_country.advertisement_id").
		Joins("left join advertisement_platform on advertisement.ID = advertisement_platform.advertisement_id")

	if query.Platform != "" {
		sqlQuery.Where(models.DB.Where("advertisement_platform.platform_name = ?", query.Platform).Or("advertisement_platform.platform_name is NULL"))
	}

	if query.Country != "" {
		sqlQuery.Where(models.DB.Where("advertisement_country.country_code = ?", query.Country).Or("advertisement_country.country_code is NULL"))
	}

	if query.Gender != "" {
		sqlQuery.Where(models.DB.Where("advertisement.Gender = ?", query.Gender).Or("advertisement.Gender is NULL"))
	}

	if query.Age != nil {
		sqlQuery.Where("advertisement.AgeStart <= ? AND advertisement.AgeEnd >= ?", query.Age, query.Age)
	}

	currentTime := time.Now()
  roundedTime := currentTime.Truncate(time.Hour)
	sqlQuery.Where("advertisement.StartAt <= ? AND advertisement.EndAt >= ?", roundedTime, roundedTime)

	sqlQuery.
		Offset(*query.Offset).
		Limit(*query.Limit).
		Order("endAt asc").
		Find(&selectResult)

	if sqlQuery.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": sqlQuery.Error})
	}

	c.JSON(http.StatusOK, gin.H{"items": selectResult})
}
