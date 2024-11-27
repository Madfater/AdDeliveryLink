package dto

import "time"

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

type GetAdsReq struct {
	Offset   *int   `form:"offset" validate:"omitempty,gte=0"`
	Limit    *int   `form:"limit" validate:"omitempty,gte=1,lte=100"`
	Age      *int   `form:"age" validate:"omitempty,gte=1,lte=100"`
	Gender   string `form:"gender" validate:"omitempty,oneof=F M"`
	Country  string `form:"country" validate:"omitempty,country_code"`
	Platform string `form:"platform" validate:"omitempty,oneof=ios android web"`
}

type GetAdsResp struct {
	Title string    `json:"Title"`
	EndAt time.Time `json:"EndAt"`
}
