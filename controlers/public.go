package controlers

import (
	"dcardAssignment/models"
	"time"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PublicControler struct{}

type Item struct {
	Title string    `json:"Title"`
	EndAt time.Time `json:"EndAt"`
}

type Result struct {
	Items []Item `json:"items"`
}

func (PublicControler) PublicAdvertisement(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	age, _ := strconv.Atoi(c.Query("age"))
	gender := c.Query("gender")
	country := c.Query("country")
	platform := c.Query("platform")

	var selectResult []models.Advertisement

	models.DB.Offset(offset).
		Limit(limit).
		Where(models.DB.Where("Conditions->'$.gender' = ?", gender).Or("JSON_EXTRACT(Conditions, '$.gender') = ''")).
		Where(models.DB.Where("JSON_CONTAINS(Conditions->'$.country', JSON_ARRAY(?)) = 1", country).Or("JSON_TYPE(Conditions->'$.country') = 'NULL'")).
		Where(models.DB.Where("JSON_CONTAINS(Conditions->'$.platform', JSON_ARRAY(?)) = 1", platform).Or("JSON_TYPE(Conditions->'$.platform') = 'NULL'")).
		Where("Conditions->'$.ageStart' <= ? and Conditions->'$.ageEnd' >= ?", age, age).
		Order("endAt asc").
		Find(&selectResult)

	var items []Item
	result := Result{Items: items}

	for _, v := range selectResult {
		result.Items = append(result.Items, Item{Title: v.Title, EndAt: v.EndAt})
	}

	c.JSON(http.StatusOK, result)
}
