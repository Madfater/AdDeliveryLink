package entity

type Country struct {
	CountryCode string `gorm:"many2many:advertisement_country;primaryKey;column:CountryCode"`
}

func (Country) TableName() string {
	return "country"
}
