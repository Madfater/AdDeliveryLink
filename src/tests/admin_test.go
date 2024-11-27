package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Madfater/AdDeliveryLink/controllers"
	"github.com/Madfater/AdDeliveryLink/dto"
	"github.com/Madfater/AdDeliveryLink/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAdvertisement(t *testing.T) {
	// 準備測試數據
	body := dto.Body{
		Title:   "Test Advertisement",
		StartAt: time.Now(),
		EndAt:   time.Now().Add(time.Hour * 24),
		Conditions: dto.Conditions{
			Country:  []string{"TW", "JP"},
			Platform: []string{"ios"},
		},
	}
	jsonBody, _ := json.Marshal(body)

	//設定MocK Server
	r := gin.Default()
	r.POST("/createAdvertisement", middleware.RequestValidator[dto.Body]{}.GetBodyValidator, controllers.AdminController{}.CreateAdvertisement)

	req, err := http.NewRequest("POST", "/createAdvertisement", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	//發送請求
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateVaildAdvertisement(t *testing.T) {
	g := "G"
	// 準備測試數據
	body := dto.Body{
		Title:   "Test Advertisement",
		StartAt: time.Now(),
		EndAt:   time.Now().Add(time.Hour * 24),
		Conditions: dto.Conditions{
			Country:  []string{"TW", "JP"},
			Platform: []string{"ios"},
			Gender:   &g,
		},
	}
	jsonBody, _ := json.Marshal(body)

	//設定MocK Server
	r := gin.Default()
	r.POST("/createAdvertisement", middleware.RequestValidator[dto.Body]{}.GetBodyValidator, controllers.AdminController{}.CreateAdvertisement)

	req, err := http.NewRequest("POST", "/createAdvertisement", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	//發送請求
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateWithEmptyConditions(t *testing.T) {
	// 準備測試數據
	body := dto.Body{
		Title:      "Sample Advertisement",
		StartAt:    time.Now(),
		EndAt:      time.Now().Add(24 * time.Hour),
		Conditions: dto.Conditions{},
	}
	jsonBody, _ := json.Marshal(body)

	//設定MocK Server
	r := gin.Default()
	r.POST("/createAdvertisement", middleware.RequestValidator[dto.Body]{}.GetBodyValidator, controllers.AdminController{}.CreateAdvertisement)

	req, err := http.NewRequest("POST", "/createAdvertisement", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	//發送請求
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
