package dto

import "time"

type Query struct {
	Offset   int   
	Limit    int
	Age      int
	Gender   string
	Country  string
	Platform string
}

type Response struct {
	Title string    `json:"Title"`
	EndAt time.Time `json:"EndAt"`
}