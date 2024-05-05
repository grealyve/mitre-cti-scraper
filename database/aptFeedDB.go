package database

import (
	"fmt"

	"github.com/grealyve/mitre-cti-scraper/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectAndMigrateDB(config *models.APTFeedConfigDB) (*gorm.DB, error) {
	// Connect to the database
	db, err := ConnectDB(config)
	if err != nil {
		return nil, err
	}

	// Using the MigrateModels function to create the tables
	if err := MigrateModels(db); err != nil {
		return nil, err
	}

	return db, nil
}

// ConnectDB is a function to connect to the database
func ConnectDB(config *models.APTFeedConfigDB) (*gorm.DB, error) {
	// Create the DSN string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return dbConnection, nil
}

// MigrateModels is a function to create the tables
func MigrateModels(db *gorm.DB) error {
	// Create the APTFeedData table
	if err := db.AutoMigrate(&models.APTFeedData{}); err != nil {
		return err
	}

	// Create the APTFeedTechnic table
	if err := db.AutoMigrate(&models.APTFeedTechnic{}); err != nil {
		return err
	}

	// Create the APTFeedTactic table
	if err := db.AutoMigrate(&models.APTFeedTactic{}); err != nil {
		return err
	}

	// Create the APTFeedRelationship table
	if err := db.AutoMigrate(&models.APTFeedRelationship{}); err != nil {
		return err
	}

	// Create the APTFeedMitigation table
	if err := db.AutoMigrate(&models.APTFeedMitigation{}); err != nil {
		return err
	}

	return nil
}

func InsertAPTFeedTechnicData(db *gorm.DB, data models.APTFeedTechnic) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
