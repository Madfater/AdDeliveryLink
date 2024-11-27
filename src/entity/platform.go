package entity

type Platform struct {
	PlatformName string `gorm:"many2many:advertisement_platform;primaryKey;column:PlatformName"`
}

func (Platform) TableName() string {
	return "platform"
}
