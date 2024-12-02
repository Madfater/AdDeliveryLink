package entity

import "github.com/Madfater/AdDeliveryLink/enum"

type Platform struct {
	PlatformName enum.Platform `gorm:"many2many:advertisement_platform;primaryKey;column:platform_name"`
}

func (Platform) TableName() string {
	return "platform"
}
