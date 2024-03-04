package database

import (
	"restapi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/btpns_task"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.UserModel{}, &models.PhotoModel{})
	DB = database
}

func GetDB() *gorm.DB {
	return DB
}
