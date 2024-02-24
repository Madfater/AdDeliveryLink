package models

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {

	//設定DSN
	user := "root"
	ip := "db"
	port := "3306"
	dbName := "DcardAssignment"

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, ip, port, dbName)

	//連線資料庫
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("Fail to connect to database: ", err)
	}
}

func CloseDB() {

	db, err := DB.DB()
	if err != nil {
		log.Fatal("Fail to get database: ", err)
	}
	db.Close()
}
