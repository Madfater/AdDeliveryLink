package dto

import (
	"github.com/Madfater/AdDeliveryLink/enum"
	"time"
)

type Filter struct {
	Platform enum.Platform
	Country  enum.CountryCode
	Gender   enum.Gender
	Age      *int
	Time     time.Time
}
