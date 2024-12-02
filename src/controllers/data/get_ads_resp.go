package data

import (
	"time"
)

type GetAdsResp struct {
	Ads []GetAdsRespItem `json:"ads"`
}

type GetAdsRespItem struct {
	Title string    `json:"Title"`
	EndAt time.Time `json:"EndAt"`
}
