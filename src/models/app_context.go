package models

import (
	"github.com/Madfater/AdDeliveryLink/models/cache"
	"github.com/Madfater/AdDeliveryLink/models/database"
	"gorm.io/gorm"
	"log"
)

type AppContext struct {
	DB *gorm.DB
}

func CreateAppContext(dbType, cacheType string) (*AppContext, error) {
	databaseFactory := &database.DatabaseFactory{}

	db, err := databaseFactory.CreateDatabase(dbType)
	if err != nil {
		return nil, err
	}

	gormDB, err := db.Connect()
	if err != nil {
		return nil, err
	}

	cacheFactory := &cache.CacheFactory{}
	cachePlugin, err := cacheFactory.CreateCachePlugin(cacheType)
	if err != nil {
		return nil, err
	}

	err = gormDB.Use(cachePlugin.GetCachePlugin())
	if err != nil {
		return nil, err
	}

	return &AppContext{DB: gormDB}, nil
}

func (ctx *AppContext) Close() {
	db, err := ctx.DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
