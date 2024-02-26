package controlers_test

import (
	"bytes"
	"dcardAssignment/controlers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAdvertisement(t *testing.T) {

	body := controlers.Body{
		Title: "Sample Advertisement",
		StartAt: time.Now(),
		EndAt: time.Now().Add(24 * time.Hour), // 结束时间为当前时间后一天
		Conditions: struct {
			AgeStart int      `json:"ageStart"`
			AgeEnd   int      `json:"ageEnd"`
			Country  []string `json:"country"`
			Platform []string `json:"platform"`
			Gender   *string  `json:"gender"`
		}{
			AgeStart: 18,
			AgeEnd:   50,
			Country:  []string{"TW", "JP"},
			Platform: []string{"android", "ios"},
			Gender:   nil, 
		},
	}
	jsonBody, _ := json.Marshal(body)

	r := gin.Default()
	
	r.POST("/createAdvertisement", controlers.AdminControler{}.CreateAdvertisement)

	req, err := http.NewRequest("POST", "/createAdvertisement", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	fmt.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "建立成功", response["Message"])
}
