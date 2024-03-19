package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Log *Logger
)

func Init() error {
	var err error

	DB, err = InicializePg()

	if err != nil {
		return fmt.Errorf("database initialization error: %v", err)
	}
	return nil
}

func GetPg() *gorm.DB {
	return DB
}

func GetLogger(p string) *Logger {
	Log = NewLogger(p)
	return Log
}
