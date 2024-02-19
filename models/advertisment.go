package models

import (
	"time"
	"gorm.io/datatypes"
)

type Advertisement struct {
	ID         uint      `gorm:"primaryKey"`
	Title      string    `gorm:"not null"`
	StartAt    time.Time `gorm:"not null"`
	EndAt      time.Time `gorm:"not null"`
	Conditions datatypes.JSON   `gorm:"not null"`
}


func (Advertisement) TableName() string {
	return "advertisement"
}


func GetAdList()(adlist []Advertisement){
	DB.Find(&adlist)
	return
}