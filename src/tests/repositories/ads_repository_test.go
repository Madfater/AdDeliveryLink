package repositories

import (
	"github.com/Madfater/AdDeliveryLink/dto"
	"github.com/Madfater/AdDeliveryLink/entity"
	"github.com/Madfater/AdDeliveryLink/enum"
	"github.com/Madfater/AdDeliveryLink/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestFindByCondition(t *testing.T) {
	// Set up in-memory SQLite database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to SQLite: %s", err)
	}

	// Migrate schema
	err = db.AutoMigrate(&entity.Advertisement{}, &entity.Country{}, &entity.Platform{})
	if err != nil {
		t.Fatalf("Failed to migrate schema: %s", err)
	}

	// Seed data
	seedTestData(db)

	// Initialize repository
	repo := repositories.NewAdsRepository(db)

	age := 25
	// Define filter
	filter := dto.Filter{
		Gender:   enum.Male,
		Platform: enum.IOS,
		Country:  enum.US,
		Age:      &age,
		Time:     time.Now(),
	}

	// Call repository method
	results, err := repo.FindByCondition(filter, 10, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	// Assert results
	expectedCount := 1
	if len(results) != expectedCount {
		t.Errorf("Expected %d results, got %d", expectedCount, len(results))
	}

	expectedTitle := "Test Ad 1"
	if len(results) > 0 && results[0].Title != expectedTitle {
		t.Errorf("Expected title %s, got %s", expectedTitle, results[0].Title)
	}
}

func seedTestData(db *gorm.DB) {
	// Add advertisements
	ad1 := entity.Advertisement{
		ID:       1,
		Title:    "Test Ad 1",
		Gender:   enum.Male,
		AgeStart: 18,
		AgeEnd:   30,
		StartAt:  time.Now().Add(-time.Hour),
		EndAt:    time.Now().Add(time.Hour),
		Status:   true,
	}
	ad2 := entity.Advertisement{
		ID:       2,
		Title:    "Test Ad 2",
		Gender:   enum.Female,
		AgeStart: 18,
		AgeEnd:   30,
		StartAt:  time.Now().Add(-time.Hour),
		EndAt:    time.Now().Add(time.Hour),
		Status:   true,
	}

	db.Create(&ad1)
	db.Create(&ad2)

	// Add related entities (Country and Platform)
	country := entity.Country{
		CountryCode: enum.US,
	}
	platform := entity.Platform{
		PlatformName: enum.IOS,
	}

	db.Create(&country)
	db.Create(&platform)

	// Add associations
	db.Exec("INSERT INTO advertisement_country (advertisement_id, country_code) VALUES (?, ?)", ad1.ID, country.CountryCode)
	db.Exec("INSERT INTO advertisement_platform (advertisement_id, platform_name) VALUES (?, ?)", ad1.ID, platform.PlatformName)
}
