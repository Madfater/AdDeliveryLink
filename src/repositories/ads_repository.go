package repositories

import (
	"github.com/Madfater/AdDeliveryLink/dto"
	"github.com/Madfater/AdDeliveryLink/entity"
	"github.com/Madfater/AdDeliveryLink/enum"
	"gorm.io/gorm"
	"time"
)

type AdsRepository interface {
	RepositoryInterface[entity.Advertisement]
	FindByCondition(filter dto.Filter, limit, offset int) ([]entity.Advertisement, error)
	ExpireAdvertisements() error
}

type adsRepository struct {
	*BaseRepository[entity.Advertisement]
}

func NewAdsRepository(db *gorm.DB) AdsRepository {
	return &adsRepository{
		BaseRepository: NewBaseRepository[entity.Advertisement](db),
	}
}

func (r *adsRepository) FindByCondition(filter dto.Filter, limit, offset int) ([]entity.Advertisement, error) {
	var results []entity.Advertisement

	query := r.db.Model(&entity.Advertisement{}).
		Distinct("advertisement.title", "advertisement.end_at").
		Preload("Country").
		Preload("Platform").
		Joins("INNER JOIN advertisement_country ON advertisement.id = advertisement_country.advertisement_id").
		Joins("INNER JOIN advertisement_platform ON advertisement.id = advertisement_platform.advertisement_id")

	query.Where("advertisement.status = ?", true)

	gender := filter.Gender
	query.Where("advertisement.gender IN (?,?)", gender, enum.Both)

	if platform := filter.Platform; platform != "" {
		query.Where("advertisement_platform.platform_name = ?", platform)
	}

	if country := filter.Country; country != "" {
		query.Where("advertisement_country.country_code = ?", country)
	}

	if age := filter.Age; age != nil {
		query.Where("? BETWEEN advertisement.age_start AND advertisement.age_end", age)
	}

	query.Where("? BETWEEN advertisement.start_at AND advertisement.end_at", filter.Time)

	err := query.Limit(limit).Offset(offset).Order("advertisement.end_at ASC").Find(&results).Error
	return results, err
}

func (r *adsRepository) ExpireAdvertisements() error {
	currentTime := time.Now().Truncate(time.Hour)
	err := r.db.Model(&entity.Advertisement{}).Where("end_at < ?", currentTime).Update("status", false).Error
	return err
}
