package config

import (
	"log"
	"os"
	"qr-menu-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var Database DBInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("qrmenu.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database. \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	db.AutoMigrate(&models.Table{}, &models.Item{}, &models.Order{}, &models.Category{})

	Database = DBInstance{Db: db}
}
