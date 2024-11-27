package services

import (
	"errors"
	"github.com/Madfater/AdDeliveryLink/dto"
	"github.com/Madfater/AdDeliveryLink/entity"
	"github.com/Madfater/AdDeliveryLink/repositories"
	"time"
)

type AdsService interface {
	CreateAdvertisement(req dto.CreateAdsReq) error
	GetAdvertisements(query dto.GetAdsReq) ([]dto.GetAdsResp, error)
}

type adsService struct {
	repo repositories.AdsRepository
}

func NewAdsService(repo repositories.AdsRepository) AdsService {
	return &adsService{repo: repo}
}

func (s *adsService) CreateAdvertisement(req dto.CreateAdsReq) error {
	defaultAgeStart := 1
	defaultAgeEnd := 100
	defaultGender := "F"

	if req.Conditions.AgeStart == nil {
		req.Conditions.AgeStart = &defaultAgeStart
	}
	if req.Conditions.AgeEnd == nil {
		req.Conditions.AgeEnd = &defaultAgeEnd
	}
	if req.Conditions.Gender == nil {
		req.Conditions.Gender = &defaultGender
	}

	ad := entity.Advertisement{
		Title:    req.Title,
		StartAt:  req.StartAt,
		EndAt:    req.EndAt,
		AgeStart: *req.Conditions.AgeStart,
		AgeEnd:   *req.Conditions.AgeEnd,
		Gender:   *req.Conditions.Gender,
		Country:  convertToCountries(req.Conditions.Country),
		Platform: convertToPlatforms(req.Conditions.Platform),
	}

	if err := s.repo.Create(&ad); err != nil {
		return errors.New("failed to create advertisement: " + err.Error())
	}
	return nil
}

func (s *adsService) GetAdvertisements(query dto.GetAdsReq) ([]dto.GetAdsResp, error) {
	// 設置預設值
	defaultLimit := 5
	defaultOffset := 0
	if query.Limit == nil {
		query.Limit = &defaultLimit
	}
	if query.Offset == nil {
		query.Offset = &defaultOffset
	}

	// 處理篩選條件
	currentTime := time.Now().Truncate(time.Hour)
	filters := map[string]interface{}{
		"platform":   query.Platform,
		"country":    query.Country,
		"gender":     query.Gender,
		"age":        query.Age,
		"start_time": currentTime,
	}

	// 調用 Repository
	results, err := s.repo.FindByCondition(filters, *query.Limit, *query.Offset)
	if err != nil {
		return nil, err
	}

	// 將實體轉換為回應格式
	var responses []dto.GetAdsResp
	for _, result := range results {
		responses = append(responses, dto.GetAdsResp{
			Title: result.Title,
			EndAt: result.EndAt,
		})
	}
	return responses, nil
}

func convertToPlatforms(platformNames []string) []entity.Platform {
	var platforms []entity.Platform
	for _, name := range platformNames {
		platforms = append(platforms, entity.Platform{PlatformName: name})
	}
	return platforms
}

func convertToCountries(platformNames []string) []entity.Country {
	var countries []entity.Country
	for _, name := range platformNames {
		countries = append(countries, entity.Country{CountryCode: name})
	}
	return countries
}
