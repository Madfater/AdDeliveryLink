package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-gorm/caches/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {

	var user string = os.Getenv("MYSQL_USERNAME")
	var ip string = os.Getenv("MYSQL_IP")
	var port string = os.Getenv("MYSQL_PORT")
	var dbName string = os.Getenv("MYSQL_Database")

	cachesPlugin := &caches.Caches{Conf: &caches.Config{
		Cacher: &redisCacher{
			rdb: redis.NewClient(&redis.Options{
				Addr:     "redis:6379",
				Password: "",
				DB:       0,
			}),
		},
	}}

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, ip, port, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal("Fail to connect to database: ", err)
	}

   _ = DB.Use(cachesPlugin)

	db, err := DB.DB()
	if err != nil {
		log.Fatal("Fail to get database: ", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
}

func CloseDB() {

	db, err := DB.DB()
	if err != nil {
		log.Fatal("Fail to get database: ", err)
	}
	db.Close()
}
