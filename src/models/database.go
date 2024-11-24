package models

import (
	"github.com/Madfater/AdDeliveryLink/models/cache"
	"github.com/Madfater/AdDeliveryLink/models/database"
	"github.com/Madfater/AdDeliveryLink/utils"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	const errMsg = "Error connecting to database"
	databaseFactory := &database.DatabaseFactory{}

	dbType := "mysql"
	db, err := databaseFactory.CreateDatabase(dbType)
	utils.HandleError(err, errMsg)

	DB, err = db.Connect()
	utils.HandleError(err, errMsg)

	cacheFactory := &cache.CacheFactory{}

	cacheType := "redis"
	cachePlugin, err := cacheFactory.CreateCachePlugin(cacheType)
	utils.HandleError(err, errMsg)

	_ = DB.Use(cachePlugin.GetCachePlugin())
}

func CloseDB() {
	db, err := DB.DB()
	utils.HandleError(err, "Fail to get database: ")
	err = db.Close()
	utils.HandleError(err, "Fail to close database: ")
}
