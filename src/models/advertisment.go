package models

import (
	"time"
)

type Advertisement struct {
	ID       uint       `gorm:"primaryKey"`
	Title    string     `gorm:"not null"`
	StartAt  time.Time  `gorm:"not null;column:StartAt"`
	EndAt    time.Time  `gorm:"not null;column:EndAt"`
	AgeStart int        `gorm:"check:AgeStart >= 1 AND AgeStart <= 100;column:AgeStart"`
	AgeEnd   int        `gorm:"check:AgeEnd >= 1 AND AgeEnd <= 100;column:AgeEnd"`
	Gender   *string    `gorm:"type:ENUM('M', 'F')"`
	Country  []Country  `gorm:"many2many:advertisement_country;foreignKey:ID;joinForeignKey:AdvertisementID;joinReferences:CountryCode;"`
	Platform []Platform `gorm:"many2many:advertisement_platform;foreignKey:ID;joinForeignKey:AdvertisementID;joinReferences:PlatformName;"`
}

type Country struct {
	CountryCode string `gorm:"many2many:advertisement_country;primaryKey;column:CountryCode"`
}

type Platform struct {
	PlatformName string `gorm:"many2many:advertisement_platform;primaryKey;column:PlatformName"`
}

func (Advertisement) TableName() string {
	return "advertisement"
}

func (Country) TableName() string {
	return "country"
}

func (Platform) TableName() string {
	return "platform"
}
