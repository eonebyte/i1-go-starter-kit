package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabasePostgreSQL(dsn string) {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	// Auto Migrate semua model
	autoMigrateAllModels()
}
