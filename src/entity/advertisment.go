package entity

import (
	"github.com/Madfater/AdDeliveryLink/enum"
	"time"
)

type Advertisement struct {
	ID          int         `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title       string      `gorm:"column:title;NOT NULL"`
	StartAt     time.Time   `gorm:"column:start_at;NOT NULL"`
	EndAt       time.Time   `gorm:"column:end_at;NOT NULL"`
	AgeStart    int         `gorm:"column:age_start"`
	AgeEnd      int         `gorm:"column:age_end"`
	Gender      enum.Gender `gorm:"column:gender"`
	Status      bool        `gorm:"column:status;;default:true"`
	CreatedDate time.Time   `gorm:"column:created_date;default:CURRENT_TIMESTAMP"`
	UpdatedDate time.Time   `gorm:"column:updated_date;default:CURRENT_TIMESTAMP"`
	Country     []Country   `gorm:"many2many:advertisement_country;foreignKey:ID;joinForeignKey:advertisement_id;joinReferences:country_code;"`
	Platform    []Platform  `gorm:"many2many:advertisement_platform;foreignKey:ID;joinForeignKey:advertisement_id;joinReferences:platform_name;"`
}

func (Advertisement) TableName() string {
	return "advertisement"
}
