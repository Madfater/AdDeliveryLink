package controllers

import (
	"errors"
	"github.com/Madfater/AdDeliveryLink/controllers"
	"github.com/Madfater/AdDeliveryLink/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// MockAdsService 定義 Mock 的 AdsService
type MockAdsService struct {
	mock.Mock
}

func (m *MockAdsService) CreateAdvertisement(req dto.CreateAdsReq) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockAdsService) GetAdvertisements(req dto.GetAdsReq) ([]dto.GetAdsResp, error) {
	args := m.Called(req)
	return args.Get(0).([]dto.GetAdsResp), args.Error(1)
}

func TestCreateAdvertisement(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success case", func(t *testing.T) {
		// Mock Service
		mockService := new(MockAdsService)
		mockService.On("CreateAdvertisement", mock.Anything).Return(nil)

		controller := controllers.NewAdsController(mockService)

		// 測試請求
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := `{"title":"Ad Title","start_at":"2024-01-01","end_at":"2024-01-31","conditions":{"gender":"male"}}`
		c.Request = httptest.NewRequest(http.MethodPost, "/ad", strings.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")

		// 執行方法
		controller.CreateAdvertisement(c)

		// 斷言
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200 but got %d", w.Code)
		}
		mockService.AssertCalled(t, "CreateAdvertisement", mock.Anything)
	})

	t.Run("invalid request body", func(t *testing.T) {
		mockService := new(MockAdsService)
		controller := controllers.NewAdsController(mockService)

		// 測試請求
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := `{"title":123}` // Invalid JSON
		c.Request = httptest.NewRequest(http.MethodPost, "/ad", strings.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")

		controller.CreateAdvertisement(c)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400 but got %d", w.Code)
		}
	})
}

func TestGetAdvertisement(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success case", func(t *testing.T) {
		// Mock Service
		mockService := new(MockAdsService)
		mockResult := []dto.GetAdsResp{
			{Title: "Ad Title 1", EndAt: time.Date(2024, time.January, 31, 0, 0, 0, 0, time.UTC)},
			{Title: "Ad Title 2", EndAt: time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC)},
		}
		mockService.On("GetAdvertisements", mock.Anything).Return(mockResult, nil)

		controller := controllers.NewAdsController(mockService)

		// 測試請求
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = httptest.NewRequest(http.MethodGet, "/ad?gender=male", nil)

		controller.GetAdvertisement(c)

		if w.Code != http.StatusOK {
			t.Errorf("expected status 200 but got %d", w.Code)
		}
		mockService.AssertCalled(t, "GetAdvertisements", mock.Anything)
	})

	t.Run("error case", func(t *testing.T) {
		mockService := new(MockAdsService)
		mockService.On("GetAdvertisements", mock.Anything).Return([]dto.GetAdsResp{}, errors.New("database error"))

		controller := controllers.NewAdsController(mockService)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = httptest.NewRequest(http.MethodGet, "/ad?gender=male", nil)

		controller.GetAdvertisement(c)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500 but got %d", w.Code)
		}
	})
}
