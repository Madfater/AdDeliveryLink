package entity

import "github.com/Madfater/AdDeliveryLink/enum"

type Country struct {
	CountryCode enum.CountryCode `gorm:"many2many:advertisement_country;primaryKey;column:country_code"`
}

func (Country) TableName() string {
	return "country"
}
