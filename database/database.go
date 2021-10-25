package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(host string, dbName string, dbUser string, dbPass string) *gorm.DB {
	dsn := "host=" + host + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
