package config

import (
	"fmt"

	"github.com/junioralcant/api-stores-go/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InicializePg() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", GetEnvVariable("DB_HOST"), GetEnvVariable("DB_USER"), GetEnvVariable("DB_PASSWORD"), GetEnvVariable("DB_NAME"), GetEnvVariable("DB_PORT"))

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
