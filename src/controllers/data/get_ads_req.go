package data

import (
	"github.com/Madfater/AdDeliveryLink/enum"
)

type GetAdsReq struct {
	Offset   *int             `form:"offset" validate:"omitempty,gte=0"`
	Limit    *int             `form:"limit" validate:"omitempty,gte=1,lte=100"`
	Age      *int             `form:"age" validate:"omitempty,gte=1,lte=100"`
	Gender   enum.Gender      `form:"gender" validate:"omitempty" binding:"omitempty,enum"`
	Country  enum.CountryCode `form:"country" validate:"omitempty" binding:"omitempty,enum"`
	Platform enum.Platform    `form:"platform" validate:"omitempty" binding:"omitempty,enum"`
}

func (req GetAdsReq) Validate() (bool, error) {
	return true, nil
}
