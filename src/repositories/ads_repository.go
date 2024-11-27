package repositories

import (
	"github.com/Madfater/AdDeliveryLink/entity"
	"gorm.io/gorm"
)

type AdsRepository interface {
	RepositoryInterface[entity.Advertisement]
	FindByCondition(filters map[string]interface{}, limit, offset int) ([]entity.Advertisement, error)
}

type adsRepository struct {
	*BaseRepository[entity.Advertisement]
}

func NewAdsRepository(db *gorm.DB) AdsRepository {
	return &adsRepository{
		BaseRepository: NewBaseRepository[entity.Advertisement](db),
	}
}

func (r *adsRepository) FindByCondition(filters map[string]interface{}, limit, offset int) ([]entity.Advertisement, error) {
	var results []entity.Advertisement

	query := r.db.Model(&entity.Advertisement{}).
		Select("DISTINCT advertisement.Title", "advertisement.EndAt").
		Joins("left join advertisement_country on advertisement.ID = advertisement_country.advertisement_id").
		Joins("left join advertisement_platform on advertisement.ID = advertisement_platform.advertisement_id")

	if platform, ok := filters["platform"]; ok {
		query.Where(r.db.Where("advertisement_platform.platform_name = ?", platform).Or("advertisement_platform.platform_name is NULL"))
	}
	if country, ok := filters["country"]; ok {
		query.Where(r.db.Where("advertisement_country.country_code = ?", country).Or("advertisement_country.country_code is NULL"))
	}
	if gender, ok := filters["gender"]; ok {
		query.Where(r.db.Where("advertisement.Gender = ?", gender).Or("advertisement.Gender is NULL"))
	}
	if age, ok := filters["age"]; ok {
		query.Where("advertisement.AgeStart <= ? AND advertisement.AgeEnd >= ?", age, age)
	}
	if startTime, ok := filters["start_time"]; ok {
		query.Where("advertisement.StartAt <= ? AND advertisement.EndAt >= ?", startTime, startTime)
	}

	err := query.Limit(limit).Offset(offset).Order("endAt asc").Find(&results).Error
	return results, err
}
