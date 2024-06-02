package database

import (
	"errors"
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

	if err := db.AutoMigrate(&models.APTFeedDataObject{}); err != nil {
		return err
	}

	// Create the APTFeedTechnic table
	if err := db.AutoMigrate(&models.APTFeedTechnic{}); err != nil {
		return err
	}

	// Create the APTFeedTechnicObject table
	if err := db.AutoMigrate(&models.APTFeedTechnicObject{}); err != nil {
		return err
	}

	// Create the APTFeedTactic table
	if err := db.AutoMigrate(&models.APTFeedTactic{}); err != nil {
		return err
	}

	// Create the APTFeedTacticObject table
	if err := db.AutoMigrate(&models.APTFeedTacticObject{}); err != nil {
		return err
	}

	// Create the APTFeedRelationship table
	if err := db.AutoMigrate(&models.APTFeedRelationship{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.APTFeedRelationshipObject{}); err != nil {
		return err
	}

	// Create the APTFeedMitigation table
	if err := db.AutoMigrate(&models.APTFeedMitigation{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.APTFeedMitigationObject{}); err != nil {
		return err
	}

	return nil
}

func InsertAPTFeedTechnicData(db *gorm.DB, data models.APTFeedTechnic) error {
	var existingData models.APTFeedTechnic
	if err := db.Where("id = ?", data.ID).First(&existingData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	} else {
		data.ID = existingData.ID
		if err := db.Save(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertAPTFeedTechnicObject(db *gorm.DB, data models.APTFeedTechnicObject) error {
	var existingData models.APTFeedTechnicObject
	if err := db.Where("id = ?", data.ID).First(&existingData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	} else {
		data.ID = existingData.ID
		if err := db.Save(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertAPTFeedData(db *gorm.DB, data models.APTFeedData) error {
	var existingData models.APTFeedData
	if err := db.Where("id = ?", data.ID).First(&existingData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	} else {
		data.ID = existingData.ID
		if err := db.Save(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertAPTFeedDataObject(db *gorm.DB, data models.APTFeedDataObject) error {
	var existingData models.APTFeedDataObject
	if err := db.Where("id = ?", data.ID).First(&existingData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	} else {
		data.ID = existingData.ID
		if err := db.Save(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertAPTFeedTacticData(db *gorm.DB, data models.APTFeedTactic) error {
	var existingData models.APTFeedTactic
	if err := db.Where("id = ?", data.ID).First(&existingData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	} else {
		data.ID = existingData.ID
		if err := db.Save(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertAPTFeedTacticObject(db *gorm.DB, data models.APTFeedTacticObject) error {
	var existingData models.APTFeedTacticObject
	if err := db.Where("id = ?", data.ID).First(&existingData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	} else {
		data.ID = existingData.ID
		if err := db.Save(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertAPTFeedRelationshipData(db *gorm.DB, data models.APTFeedRelationship) error {
	var existingData models.APTFeedRelationship
	if err := db.Where("id = ?", data.ID).First(&existingData).Error; err != nil {
		// Veri yoksa, veriyi ekleme
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	} else {
		data.ID = existingData.ID
		if err := db.Save(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertAPTFeedRelationshipObject(db *gorm.DB, data models.APTFeedRelationshipObject) error {
	var existingData models.APTFeedRelationshipObject
	if err := db.Where("id = ?", data.ID).First(&existingData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	} else {
		data.ID = existingData.ID
		if err := db.Save(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertAPTFeedMitigationData(db *gorm.DB, data models.APTFeedMitigation) error {
	var existingData models.APTFeedMitigation
	if err := db.Where("id = ?", data.ID).First(&existingData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	} else {
		data.ID = existingData.ID
		if err := db.Save(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

func InsertAPTFeedMitigationObject(db *gorm.DB, data models.APTFeedMitigationObject) error {
	var existingData models.APTFeedMitigationObject
	if err := db.Where("id = ?", data.ID).First(&existingData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	} else {
		data.ID = existingData.ID
		if err := db.Save(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetAPTFeedTechnicObjectDataList(db *gorm.DB) ([]models.APTFeedTechnicObject, error) {
	var data []models.APTFeedTechnicObject
	if err := db.Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func GetAPTFeedDataObjectDataList(db *gorm.DB) ([]models.APTFeedDataObject, error) {
	var data []models.APTFeedDataObject
	if err := db.Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func GetAPTFeedTacticObjectDataList(db *gorm.DB) ([]models.APTFeedTacticObject, error) {
	var data []models.APTFeedTacticObject
	if err := db.Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func GetAPTFeedRelationshipObjectDataList(db *gorm.DB) ([]models.APTFeedRelationshipObject, error) {
	var data []models.APTFeedRelationshipObject
	if err := db.Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func GetAPTFeedMitigationObjectDataList(db *gorm.DB) ([]models.APTFeedMitigationObject, error) {
	var data []models.APTFeedMitigationObject
	if err := db.Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}