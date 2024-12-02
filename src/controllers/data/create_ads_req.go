package data

import (
	"errors"
	"github.com/Madfater/AdDeliveryLink/enum"
	"time"
)

type CreateAdsReq struct {
	Title      string        `json:"title" validate:"required"`
	StartAt    time.Time     `json:"startAt" validate:"required" example:"2021-01-01T00:00:00Z"`
	EndAt      time.Time     `json:"endAt" validate:"required" example:"2021-01-01T00:00:00Z"`
	Conditions AdsConditions `json:"conditions" validate:"required"`
}

type AdsConditions struct {
	AgeStart *int               `json:"ageStart" validate:"omitempty,gte=1,lte=100"`
	AgeEnd   *int               `json:"ageEnd" validate:"omitempty,gte=1,lte=100"`
	Country  []enum.CountryCode `json:"country" validate:"required,dive" binding:"required,enum"`
	Platform []enum.Platform    `json:"platform" validate:"required,dive" binding:"required,enum"`
	Gender   enum.Gender        `json:"gender" validate:"required" binding:"required,enum"`
}

func (req CreateAdsReq) Validate() (bool, error) {
	if req.StartAt.After(req.EndAt) {
		return false, errors.New("StartAt must be less than EndAt")
	}

	if *req.Conditions.AgeStart > *req.Conditions.AgeEnd {
		return false, errors.New("AgeStart must be less than AgeEnd")
	}

	return true, nil
}
