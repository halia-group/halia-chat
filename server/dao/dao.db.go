package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/chat?charset=utf8mb4,utf8&parseTime=true&loc=Asia%2FShanghai"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	return db, nil
}
