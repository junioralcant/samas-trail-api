package config

import (
	"github.com/junioralcant/api-stores-go/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InicializePg() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	dsn := "host=localhost user=postgres password=toor dbname=samastrail port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("postgres initialization error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		logger.Errorf("postgres auto migration error: %v", err)
		return nil, err
	}

	return db, nil
}
