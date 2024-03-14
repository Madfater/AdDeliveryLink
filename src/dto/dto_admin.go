package dto

import "time"

type Body struct {
	Title      string     `json:"title" validate:"required"`
	StartAt    time.Time  `json:"startAt" validate:"required"`
	EndAt      time.Time  `json:"endAt" validate:"required"`
	Conditions Conditions `json:"conditions" validate:"required"`
}

type Conditions struct {
	AgeStart int      `json:"ageStart" validate:"min=1,max=100"`
	AgeEnd   int      `json:"ageEnd" validate:"min=1,max=100"`
	Country  []string `json:"country"`
	Platform []string `json:"platform"`
	Gender   *string  `json:"gender" `
}