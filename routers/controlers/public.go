package controlers

import (
	"dcardAssignment/models"
	//"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PublicControler struct{}

func (PublicControler) PublicAdvertisement(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	age, _ := strconv.Atoi(c.Query("age"))
	gender := c.Query("gender")
	country := c.Query("country")
	platform := c.Query("platform")

	var res []models.Advertisement

	models.DB.Offset(offset).
		Limit(limit).
		Where("Conditions->'$.gender' = ?", gender).
		Or("JSON_EXTRACT(Conditions, '$.gender') = """).
		Where("JSON_CONTAINS(Conditions->'$.country', JSON_ARRAY(?)) = 1", country).
		Or("JSON_TYPE(Conditions->'$.country') = 'NULL'").
		Where("JSON_CONTAINS(Conditions->'$.platform', JSON_ARRAY(?)) = 1", platform).
		Or("JSON_TYPE(Conditions->'$.platform') = 'NULL'").
		Where("Conditions->'$.ageStart' <= ? and Conditions->'$.ageEnd' >= ?", age, age).
		Find(&res)

	c.JSON(http.StatusOK, gin.H{"message": res})

}
