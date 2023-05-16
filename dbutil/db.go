package dbutil

import (
	"example/restfulapi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbURLLocal = "bookrestapi:asdfasdf@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB

func connectDb(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dbURLLocal), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitDb() (*gorm.DB, error) {
	db, err := connectDb(dbURLLocal)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
