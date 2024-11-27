package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Madfater/AdDeliveryLink/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func setupTestDBWithMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func TestAdsRepository_FindByCondition(t *testing.T) {
	gormDB, mock, err := setupTestDBWithMock()
	if err != nil {
		t.Fatalf("failed to set up mock database: %s", err)
	}
	repo := repositories.NewAdsRepository(gormDB)

	t.Run("successfully find advertisements with filters", func(t *testing.T) {
		expectedSQL := `SELECT\s+DISTINCT\s+advertisement\.Title,advertisement\.EndAt\s+` +
			`FROM\s+\` + "`advertisement`" + `\s+` +
			`left\s+join\s+advertisement_country\s+on\s+advertisement\.ID\s+=\s+advertisement_country\.advertisement_id\s+` +
			`left\s+join\s+advertisement_platform\s+on\s+advertisement\.ID\s+=\s+advertisement_platform\.advertisement_id\s+` +
			`WHERE\s+` +
			`\(advertisement_platform\.platform_name\s+=\s+\?\s+OR\s+advertisement_platform\.platform_name\s+is\s+NULL\)\s+` +
			`AND\s+\(advertisement_country\.country_code\s+=\s+\?\s+OR\s+advertisement_country\.country_code\s+is\s+NULL\)\s+` +
			`AND\s+\(advertisement\.Gender\s+=\s+\?\s+OR\s+advertisement\.Gender\s+is\s+NULL\)\s+` +
			`AND\s+\(advertisement\.AgeStart\s+<=\s+\?\s+AND\s+advertisement\.AgeEnd\s+>=\s+\?\)\s+` +
			`AND\s+\(advertisement\.StartAt\s+<=\s+\?\s+AND\s+advertisement\.EndAt\s+>=\s+\?\)\s+` +
			`ORDER\s+BY\s+endAt\s+asc\s+LIMIT\s+\?`

		mock.ExpectQuery(expectedSQL).
			WithArgs("mobile", "US", "male", 25, 25, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), 10).
			WillReturnRows(sqlmock.NewRows([]string{"Title", "EndAt"}).
				AddRow("Ad 1", time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC)). // 修正類型為 time.Time
				AddRow("Ad 2", time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC)))

		filters := map[string]interface{}{
			"platform":   "mobile",
			"country":    "US",
			"gender":     "male",
			"age":        25,
			"start_time": time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		}

		results, err := repo.FindByCondition(filters, 10, 0)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		// 驗證結果
		if len(results) != 2 || results[0].Title != "Ad 1" || results[1].Title != "Ad 2" {
			t.Errorf("unexpected result: %v", results)
		}

		// 驗證是否所有期望的查詢都被觸發
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("no advertisements found", func(t *testing.T) {
		expectedSQL := `SELECT\s+DISTINCT\s+advertisement\.Title,advertisement\.EndAt\s+` +
			`FROM\s+\` + "`advertisement`" + `\s+` +
			`left\s+join\s+advertisement_country\s+on\s+advertisement\.ID\s+=\s+advertisement_country\.advertisement_id\s+` +
			`left\s+join\s+advertisement_platform\s+on\s+advertisement\.ID\s+=\s+advertisement_platform\.advertisement_id\s+` +
			`ORDER\s+BY\s+endAt\s+asc\s+LIMIT\s+\?`

		mock.ExpectQuery(expectedSQL).
			WithArgs(10).
			WillReturnRows(sqlmock.NewRows([]string{"Title", "EndAt"}))

		filters := map[string]interface{}{}

		results, err := repo.FindByCondition(filters, 10, 0)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		// 驗證結果
		if len(results) != 0 {
			t.Errorf("expected no results, got: %v", results)
		}

		// 驗證是否所有期望的查詢都被觸發
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
