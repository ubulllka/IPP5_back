package config

import (
	"example/back/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host			= "localhost"
	port			= 5432
	user			= "postgres"
	password	= "1234"
	dbName		= "ipp5"
)

var DB *gorm.DB

func DatabaseConnection() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Contact{})
	DB = db
}