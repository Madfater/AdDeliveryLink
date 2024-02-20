package controlers

import (
	"dcardAssignment/models"
	
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminControler struct{}

type Body struct {
	Title      string    `json:"title"`
	StartAt    time.Time `json:"startAt"`
	EndAt      time.Time `json:"endAt"`
	Conditions struct {
		AgeStart int      `json:"ageStart"`
		AgeEnd   int      `json:"ageEnd"`
		Country  []string `json:"country"`
		Platform []string `json:"platform"`
	} `json:"conditions"`
}

func (AdminControler) CreateAdvertisement(c *gin.Context) {
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		// 處理錯誤
		c.JSON(http.StatusBadRequest, gin.H{"Loadin Body failed": err.Error()})
		return
	}

	json, err := json.Marshal(body.Conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"JSON encoding failed:": err})
		return
	}

	ad := models.Advertisement{
		Title:      body.Title,
		StartAt:    body.StartAt,
		EndAt:      body.EndAt,
		Conditions: json,
	}

	res := models.DB.Create(&ad)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Create failed:": err})
		return
	}

	c.JSON(http.StatusOK,gin.H{"Message":"建立成功"})
}
