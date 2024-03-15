package dto

import "time"

type Body struct {
	Title      string     `json:"title" validate:"required"`
	StartAt    time.Time  `json:"startAt" validate:"required"`
	EndAt      time.Time  `json:"endAt" validate:"required"`
	Conditions Conditions `json:"conditions" validate:"required"`
}

type Conditions struct {
	AgeStart *int     `json:"ageStart" validate:"omitempty,gte=1,lte=100"`
	AgeEnd   *int     `json:"ageEnd" validate:"omitempty,gte=1,lte=100"`
	Country  []string `json:"country" validate:"omitempty,dive,country_code"`
	Platform []string `json:"platform" validate:"omitempty,dive,oneof=ios android web"`
	Gender   *string   `json:"gender" validate:"omitempty,oneof=M F"`
}
