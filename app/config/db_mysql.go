package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabaseMySQL(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	// Auto Migrate semua model
	autoMigrateAllModels()
}
