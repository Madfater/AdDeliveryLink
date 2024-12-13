package database

import (
	"fmt"
	"github.com/Madfater/AdDeliveryLink/log"
	"gorm.io/gorm/logger"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB
}

func (m *MySQL) Connect() (*gorm.DB, error) {
	user := os.Getenv("MYSQL_USERNAME")
	ip := os.Getenv("MYSQL_IP")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	_ = log.GetLogger()
	gormLogger := NewGormLogger(logger.Silent, 2*time.Minute)

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, ip, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 gormLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	m.DB = db
	return db, nil
}

func (m *MySQL) Close() error {
	sqlDB, err := m.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
