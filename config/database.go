package config

import (
	// "database/sql"
	"example/back/model"
	// "fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)




var DB *gorm.DB

func DatabaseConnection() {
	godotenv.Load()
	var (
		// host			= "localhost"
		// port			= 5432
		// user			= "postgres"
		// password	= "1234"
		// dbName		= "ipp5"
	)

	// sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	sqlInfo := os.Getenv("DB_CONNECT")
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Contact{})
	DB = db
}