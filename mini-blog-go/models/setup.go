package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=example dbname=blog port=5431 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to connect to db")
	}
	err = database.AutoMigrate(&Post{}, &User{})

	if err != nil {
		return
	}
	DB = database
}
