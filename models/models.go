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

	//讀取配置文件
	cfg, cfgErr := ini.Load("config.ini")
	if cfgErr != nil {
		log.Fatal("Fail to read file: ", cfgErr)
	}

	//設定DSN
	user := cfg.Section("mysql").Key("user").String()
	password := cfg.Section("mysql").Key("password").String()
	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	dbName := "DcardAssignment"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, ip, port, dbName)

	//連線資料庫
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
