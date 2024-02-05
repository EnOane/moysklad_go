package infr

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moysklad/config"
	"moysklad/infr/models"
)

var Db *gorm.DB

func ConnectToDb() {
	dbConfig := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.AppConfig.PostgresConfig.Host,
		config.AppConfig.PostgresConfig.Port,
		config.AppConfig.PostgresConfig.User,
		config.AppConfig.PostgresConfig.Password,
		config.AppConfig.PostgresConfig.DbName)

	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Db = db

	// Migrate the schemas
	err = db.AutoMigrate(&models.AssortmentModel{})
	if err != nil {
		return
	}

	fmt.Println("Successfully connected to database!")
}
