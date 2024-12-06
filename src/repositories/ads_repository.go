package repositories

import (
	"github.com/Madfater/AdDeliveryLink/dto"
	"github.com/Madfater/AdDeliveryLink/entity"
	"gorm.io/gorm"
)

type AdsRepository interface {
	RepositoryInterface[entity.Advertisement]
	FindByCondition(filter dto.Filter, limit, offset int) ([]entity.Advertisement, error)
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
		Select("DISTINCT advertisement.Title", "advertisement.end_at").
		Joins("left join advertisement_country on advertisement.ID = advertisement_country.advertisement_id").
		Joins("left join advertisement_platform on advertisement.ID = advertisement_platform.advertisement_id")

	query.Where("advertisement.status = ?", true)

	if gender := filter.Gender; gender != "" {
		query.Where(r.db.Where("advertisement.Gender = ?", gender).Or("advertisement.Gender = B"))
	}

	if platform := filter.Platform; platform != "" {
		query.Where(r.db.Where("advertisement_platform.platform_name = ?", platform))
	}

	if country := filter.Country; country != "" {
		query.Where(r.db.Where("advertisement_country.country_code = ?", country))
	}

	if age := filter.Age; age != nil {
		query.Where("advertisement.age_start <= ? AND advertisement.age_end >= ?", age, age)
	}

	query.Where("advertisement.start_at <= ? AND advertisement.end_at >= ?", filter.Time, filter.Time)

	err := query.Limit(limit).Offset(offset).Order("advertisement.end_at ASC").Find(&results).Error
	return results, err
}
