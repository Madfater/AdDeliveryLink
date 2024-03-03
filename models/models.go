package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {

	//確定環境
	mode := os.Getenv("GIN_MODE")
	var user string
	var ip string
	var port string
	var dbName string

	if mode == "release" {
		user = "root"
		ip = "db"
		port = "3306"
		dbName = "DcardAssignment"
	} else {
		user = "DcardAssignment:!Dcard0219"
		ip = "localhost"
		port = "3306"
		dbName = "DcardAssignment"
	}

	//設定DSN
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, ip, port, dbName)

	//連線資料庫
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal("Fail to connect to database: ", err)
	}

	db, err := DB.DB()
	if err != nil {
		log.Fatal("Fail to get database: ", err)
	}

	//資料庫設定以提升效能
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