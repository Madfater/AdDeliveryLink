package services

import (
	"github.com/Madfater/AdDeliveryLink/controllers/data"
	"github.com/Madfater/AdDeliveryLink/dto"
	"github.com/Madfater/AdDeliveryLink/entity"
	"github.com/Madfater/AdDeliveryLink/enum"
	"github.com/Madfater/AdDeliveryLink/repositories"
	"github.com/Madfater/AdDeliveryLink/utils"
	"time"
)

type AdsService interface {
	CreateAdvertisement(req data.CreateAdsReq) (data.GenericResponse[entity.Advertisement], error)
	GetAdvertisements(query data.GetAdsReq) (data.GenericResponse[data.GetAdsResp], error)
}

type adsService struct {
	repo repositories.AdsRepository
}

func NewAdsService(repo repositories.AdsRepository) AdsService {
	return &adsService{repo: repo}
}

func (s *adsService) CreateAdvertisement(req data.CreateAdsReq) (data.GenericResponse[entity.Advertisement], error) {
	ageStart := 1
	ageEnd := 100
	gender := enum.Gender(`B`)

	if req.Conditions.AgeStart != nil {
		ageStart = *req.Conditions.AgeStart
	}
	if req.Conditions.AgeEnd != nil {
		ageEnd = *req.Conditions.AgeEnd
	}
	if req.Conditions.Gender != "" {
		gender = req.Conditions.Gender
	}

	ad := entity.Advertisement{
		Title:    req.Title,
		StartAt:  req.StartAt,
		EndAt:    req.EndAt,
		AgeStart: ageStart,
		AgeEnd:   ageEnd,
		Gender:   gender,
		Country: utils.SliceConvertor(req.Conditions.Country, func(name enum.CountryCode) entity.Country {
			return entity.Country{CountryCode: name}
		}),
		Platform: utils.SliceConvertor(req.Conditions.Platform, func(name enum.Platform) entity.Platform {
			return entity.Platform{PlatformName: name}
		}),
	}

	if err := s.repo.Create(&ad); err != nil {
		return data.GenericResponse[entity.Advertisement]{}, err
	}

	return data.GenericResponse[entity.Advertisement]{
		Status:  "success",
		Message: "Ad create successfully",
		Result:  ad,
	}, nil
}

func (s *adsService) GetAdvertisements(query data.GetAdsReq) (data.GenericResponse[data.GetAdsResp], error) {
	limit := 5
	offset := 0
	if query.Limit != nil {
		limit = *query.Limit
	}
	if query.Offset != nil {
		offset = *query.Offset
	}

	currentTime := time.Now().Truncate(time.Hour)
	filters := dto.Filter{
		Platform: query.Platform,
		Country:  query.Country,
		Gender:   query.Gender,
		Age:      query.Age,
		Time:     currentTime,
	}

	results, err := s.repo.FindByCondition(filters, limit, offset)
	if err != nil {
		return data.GenericResponse[data.GetAdsResp]{}, err
	}

	responseData := utils.SliceConvertor(results, func(result entity.Advertisement) data.GetAdsRespItem {
		return data.GetAdsRespItem{
			Title: result.Title,
			EndAt: result.EndAt,
		}
	})

	return data.GenericResponse[data.GetAdsResp]{
		Status:  "success",
		Message: "Ads fetched successfully",
		Result:  data.GetAdsResp{Ads: responseData},
	}, nil
}
