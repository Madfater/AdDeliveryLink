package controllers

import (
	"errors"
	"github.com/Madfater/AdDeliveryLink/controllers"
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"github.com/Madfater/AdDeliveryLink/entity"
	"github.com/Madfater/AdDeliveryLink/middleware"
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

func (m *MockAdsService) CreateAdvertisement(req data.CreateAdsReq) (data.IResponse[entity.Advertisement], error) {
	args := m.Called(req)

	// 使用 args.Get(0) 確保回傳的泛型類型正確
	response, _ := args.Get(0).(data.IResponse[entity.Advertisement])

	return response, args.Error(1)
}

func (m *MockAdsService) GetAdvertisements(query data.GetAdsReq) (data.IResponse[data.GetAdsResp], error) {
	args := m.Called(query)
	return args.Get(0).(data.IResponse[data.GetAdsResp]), args.Error(1)
}

func TestCreateAdvertisement(t *testing.T) {
	gin.SetMode(gin.TestMode)

	middleware.RegisterCustomValidation()

	t.Run("success case", func(t *testing.T) {
		// Mock Service
		mockService := new(MockAdsService)
		mockResponse := data.IResponse[entity.Advertisement]{
			Status:  "success",
			Message: "Advertisement created successfully",
			Result: entity.Advertisement{
				Title:   "Ad Title",
				StartAt: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndAt:   time.Date(2024, time.January, 31, 0, 0, 0, 0, time.UTC),
			},
		}
		mockService.On("CreateAdvertisement", mock.Anything).Return(mockResponse, nil)

		controller := controllers.NewAdsController(mockService)

		// 測試請求
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := `{"title":"Ad Title","start_at":"2024-01-01","end_at":"2024-01-31","conditions":{"gender":"M","age_start":1,"age_end":100,"country":["US"],"platform":["ios"]}}`
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

		// 斷言
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected status 400 but got %d", w.Code)
		}
	})
}

func TestGetAdvertisement(t *testing.T) {
	gin.SetMode(gin.TestMode)

	middleware.RegisterCustomValidation()

	t.Run("success case", func(t *testing.T) {
		// Mock Service
		mockService := new(MockAdsService)
		mockResponse := data.IResponse[data.GetAdsResp]{
			Status:  "success",
			Message: "Advertisements fetched successfully",
			Result: data.GetAdsResp{
				Ads: []data.GetAdsRespItem{
					{Title: "Ad Title 1", EndAt: time.Date(2024, time.January, 31, 0, 0, 0, 0, time.UTC)},
					{Title: "Ad Title 2", EndAt: time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC)},
				},
			},
		}
		mockService.On("GetAdvertisements", mock.Anything).Return(mockResponse, nil)

		controller := controllers.NewAdsController(mockService)

		// 測試請求
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = httptest.NewRequest(http.MethodGet, "/ad?gender=F", nil)

		controller.GetAdvertisement(c)

		// 斷言
		if w.Code != http.StatusOK {
			t.Errorf("expected status 200 but got %d", w.Code)
		}
		mockService.AssertCalled(t, "GetAdvertisements", mock.Anything)
	})

	t.Run("error case", func(t *testing.T) {
		mockService := new(MockAdsService)
		mockResponse := data.IResponse[data.GetAdsResp]{
			Status:  "error",
			Message: "Failed to fetch advertisements",
			Result:  data.GetAdsResp{},
		}
		mockService.On("GetAdvertisements", mock.Anything).Return(mockResponse, errors.New("database error"))

		controller := controllers.NewAdsController(mockService)

		// 測試請求
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = httptest.NewRequest(http.MethodGet, "/ad?gender=F", nil)

		controller.GetAdvertisement(c)

		// 斷言
		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected status 500 but got %d", w.Code)
		}
	})
}
