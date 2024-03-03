package controlers_test

import (
	"bytes"
	"dcardAssignment/controlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateAdvertisement(t *testing.T) {
	// 準備測試數據
	body := controlers.Body{
		Title:   "Test Advertisement",
		StartAt: time.Now(),
		EndAt:   time.Now().Add(time.Hour * 24),
		Conditions: controlers.Conditions{
			Country:  []string{"TW", "JP"},
			Platform: []string{"ios"},
			Gender:   nil,
		},
	}
	jsonBody, _ := json.Marshal(body)

	//設定MocK Server
	r := gin.Default()
	r.POST("/createAdvertisement", controlers.AdminControler{}.CreateAdvertisement)

	req, err := http.NewRequest("POST", "/createAdvertisement", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	//發送請求
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "建立成功", response["Message"])
}


func TestCreateVaildAdvertisement(t *testing.T) {
	// 準備測試數據
	gender := "G"
	body := controlers.Body{
		Title:   "Test Advertisement",
		StartAt: time.Now(),
		EndAt:   time.Now().Add(time.Hour * 24),
		Conditions: controlers.Conditions{
			Country:  []string{"TW", "JP"},
			Platform: []string{"ios"},
			Gender:   &gender,
		},
	}
	jsonBody, _ := json.Marshal(body)

	//設定MocK Server
	r := gin.Default()
	r.POST("/createAdvertisement", controlers.AdminControler{}.CreateAdvertisement)

	req, err := http.NewRequest("POST", "/createAdvertisement", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	//發送請求
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, response["Create failed:"])
}

func TestCreateWithEmptyConditions(t *testing.T) {
	// 準備測試數據
	body := controlers.Body{
		Title:   "Sample Advertisement",
		StartAt: time.Now(),
		EndAt:   time.Now().Add(24 * time.Hour),
		Conditions: controlers.Conditions{},
	}
	jsonBody, _ := json.Marshal(body)

	//設定MocK Server
	r := gin.Default()
	r.POST("/createAdvertisement", controlers.AdminControler{}.CreateAdvertisement)

	req, err := http.NewRequest("POST", "/createAdvertisement", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	//發送請求
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "建立成功", response["Message"])
}
