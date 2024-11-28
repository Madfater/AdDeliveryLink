package data

import (
	"errors"
	"time"
)

type CreateAdsReq struct {
	Title      string        `json:"title" validate:"required"`
	StartAt    time.Time     `json:"startAt" validate:"required"`
	EndAt      time.Time     `json:"endAt" validate:"required"`
	Conditions AdsConditions `json:"conditions" validate:"required"`
}

type AdsConditions struct {
	AgeStart *int     `json:"ageStart" validate:"omitempty,gte=1,lte=100"`
	AgeEnd   *int     `json:"ageEnd" validate:"omitempty,gte=1,lte=100"`
	Country  []string `json:"country" validate:"omitempty,dive,country_code"`
	Platform []string `json:"platform" validate:"omitempty,dive,oneof=ios android web"`
	Gender   *string  `json:"gender" validate:"omitempty,oneof=M F"`
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
