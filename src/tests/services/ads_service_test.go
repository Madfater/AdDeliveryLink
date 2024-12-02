package services

import (
	"errors"
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"github.com/Madfater/AdDeliveryLink/dto"
	"github.com/Madfater/AdDeliveryLink/entity"
	"github.com/Madfater/AdDeliveryLink/enum"
	"github.com/Madfater/AdDeliveryLink/services"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// MockAdsRepository 定義 Mock 的 AdsRepository
type MockAdsRepository struct {
	mock.Mock
}

func (m *MockAdsRepository) FindByID(id uint) (*entity.Advertisement, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Advertisement), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAdsRepository) Update(ad *entity.Advertisement) error {
	args := m.Called(ad)
	return args.Error(0)
}

func (m *MockAdsRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAdsRepository) Create(ad *entity.Advertisement) error {
	args := m.Called(ad)
	return args.Error(0)
}

func (m *MockAdsRepository) FindByCondition(filters dto.Filter, limit, offset int) ([]entity.Advertisement, error) {
	args := m.Called(filters, limit, offset)
	return args.Get(0).([]entity.Advertisement), args.Error(1)
}

func TestCreateAdvertisement(t *testing.T) {

	t.Run("successfully create advertisement", func(t *testing.T) {
		mockRepo := new(MockAdsRepository)
		service := services.NewAdsService(mockRepo)

		req := data.CreateAdsReq{
			Title:   "Test Ad",
			StartAt: time.Now(),
			EndAt:   time.Now().Add(24 * time.Hour),
			Conditions: data.AdsConditions{
				Gender:   "F",
				Platform: []enum.Platform{enum.IOS},
				Country:  []enum.CountryCode{enum.US},
			},
		}

		// Mock Repository 的 Create 方法
		mockRepo.On("Create", mock.Anything).Return(nil)

		// 調用 Service
		err := service.CreateAdvertisement(req)

		// 驗證行為
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		mockRepo.AssertCalled(t, "Create", mock.MatchedBy(func(ad *entity.Advertisement) bool {
			return ad.Title == "Test Ad" && ad.AgeStart == 1 && ad.AgeEnd == 100 && ad.Gender == "F"
		}))
	})

	t.Run("failed to create advertisement", func(t *testing.T) {
		mockRepo := new(MockAdsRepository)
		service := services.NewAdsService(mockRepo)

		req := data.CreateAdsReq{
			Title:   "Test Ad",
			StartAt: time.Now(),
			EndAt:   time.Now().Add(24 * time.Hour),
		}

		// 模擬 Repository 回傳錯誤
		mockRepo.On("Create", mock.Anything).Return(errors.New("repository error"))

		// 調用 Service
		err := service.CreateAdvertisement(req)

		// 驗證行為
		if err == nil || err.Error() != "failed to create advertisement: repository error" {
			t.Errorf("expected repository error, got %v", err)
		}
	})
}

func TestGetAdvertisements(t *testing.T) {
	t.Run("successfully fetch advertisements", func(t *testing.T) {
		mockRepo := new(MockAdsRepository)
		service := services.NewAdsService(mockRepo)

		query := data.GetAdsReq{
			Platform: "mobile",
			Country:  "US",
			Age:      new(int),
			Limit:    nil, // 測試預設值
			Offset:   nil,
		}
		*query.Age = 25
		mockResult := []entity.Advertisement{
			{Title: "Ad 1", EndAt: time.Now().Add(24 * time.Hour)},
			{Title: "Ad 2", EndAt: time.Now().Add(48 * time.Hour)},
		}

		// Mock Repository 的 FindByCondition 方法
		mockRepo.On("FindByCondition", mock.Anything, 5, 0).Return(mockResult, nil)

		// 調用 Service
		results, err := service.GetAdvertisements(query)

		// 驗證行為
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(results) != 2 || results[0].Title != "Ad 1" {
			t.Errorf("unexpected results: %v", results)
		}

	})

	t.Run("failed to fetch advertisements", func(t *testing.T) {
		mockRepo := new(MockAdsRepository)
		service := services.NewAdsService(mockRepo)

		query := data.GetAdsReq{}

		mockRepo.On("FindByCondition", mock.Anything, 5, 0).Return([]entity.Advertisement{}, errors.New("repository error"))

		// 調用 Service
		results, err := service.GetAdvertisements(query)

		// 驗證行為
		if err == nil || err.Error() != "repository error" {
			t.Errorf("expected repository error, got %v", err)
		}
		if results != nil {
			t.Errorf("expected nil results, got %v", results)
		}
	})
}
