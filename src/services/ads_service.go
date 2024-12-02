package services

import (
	"errors"
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"github.com/Madfater/AdDeliveryLink/dto"
	"github.com/Madfater/AdDeliveryLink/entity"
	"github.com/Madfater/AdDeliveryLink/enum"
	"github.com/Madfater/AdDeliveryLink/repositories"
	"time"
)

type AdsService interface {
	CreateAdvertisement(req data.CreateAdsReq) error
	GetAdvertisements(query data.GetAdsReq) ([]data.GetAdsResp, error)
}

type adsService struct {
	repo repositories.AdsRepository
}

func NewAdsService(repo repositories.AdsRepository) AdsService {
	return &adsService{repo: repo}
}

func (s *adsService) CreateAdvertisement(req data.CreateAdsReq) error {
	ageStart := 1
	ageEnd := 100

	if req.Conditions.AgeStart != nil {
		ageStart = *req.Conditions.AgeStart
	}
	if req.Conditions.AgeEnd != nil {
		ageEnd = *req.Conditions.AgeEnd
	}

	ad := entity.Advertisement{
		Title:    req.Title,
		StartAt:  req.StartAt,
		EndAt:    req.EndAt,
		AgeStart: ageStart,
		AgeEnd:   ageEnd,
		Gender:   req.Conditions.Gender,
		Country:  convertToCountries(req.Conditions.Country),
		Platform: convertToPlatforms(req.Conditions.Platform),
	}

	if err := s.repo.Create(&ad); err != nil {
		return errors.New("failed to create advertisement: " + err.Error())
	}
	return nil
}

func (s *adsService) GetAdvertisements(query data.GetAdsReq) ([]data.GetAdsResp, error) {
	// 設置預設值
	limit := 5
	offset := 0
	if query.Limit != nil {
		limit = *query.Limit
	}
	if query.Offset != nil {
		offset = *query.Offset
	}

	// 處理篩選條件
	currentTime := time.Now().Truncate(time.Hour)
	filters := dto.Filter{
		Platform: query.Platform,
		Country:  query.Country,
		Gender:   query.Gender,
		Age:      query.Age,
		Time:     currentTime,
	}

	// 調用 Repository
	results, err := s.repo.FindByCondition(filters, limit, offset)
	if err != nil {
		return nil, err
	}

	// 將實體轉換為回應格式
	var responses []data.GetAdsResp
	for _, result := range results {
		responses = append(responses, data.GetAdsResp{
			Title: result.Title,
			EndAt: result.EndAt,
		})
	}
	return responses, nil
}

func convertToPlatforms(platformNames []enum.Platform) []entity.Platform {
	var platforms []entity.Platform
	for _, name := range platformNames {
		platforms = append(platforms, entity.Platform{PlatformName: name})
	}
	return platforms
}

func convertToCountries(platformNames []enum.CountryCode) []entity.Country {
	var countries []entity.Country
	for _, name := range platformNames {
		countries = append(countries, entity.Country{CountryCode: name})
	}
	return countries
}
