package data

import "time"

type GetAdsResp struct {
	Title string    `json:"Title"`
	EndAt time.Time `json:"EndAt"`
}
