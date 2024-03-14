package dto

import "time"

type Query struct {
	Offset   int `form:"offset" validate:"required"`
	Limit    int `form:"limit" validate:"required"`
	Age      int `form:"age"`
	Gender   string `form:"gender"`
	Country  string `form:"country"`
	Platform string `form:"platform"`
}

type Response struct {
	Title string    `json:"Title"`
	EndAt time.Time `json:"EndAt"`
}
