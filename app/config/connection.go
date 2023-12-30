package config

import (
	"i1-go-starter-kit/app/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func InitEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var DB *gorm.DB

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func GetDatabaseConfig() DatabaseConfig {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return DatabaseConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUser,
		Password: dbPassword,
		DBName:   dbName,
	}
}

func (c *DatabaseConfig) GetDSN(dsnType string) string {
	switch dsnType {
	case "mysql":
		return c.Username + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	case "postgresql":
		return "host=" + c.Host + " user=" + c.Username + " password=" + c.Password + " dbname=" + c.DBName + " port=" + c.Port + "sslmode=disable TimeZone=Asia/Shanghai"
	default:
		return ""
	}
}

func autoMigrateAllModels() {
	// Ambil semua model yang ada dalam package GORM
	models := []interface{}{
		&models.User{},
		// Tambahkan model lain jika ada
		// &ModelLain{},
	}

	for _, model := range models {
		if !DB.Migrator().HasTable(model) {
			// Lakukan migrasi jika tabel belum ada
			if err := DB.AutoMigrate(model); err != nil {
				panic("failed to migrate table")
			}
		}
	}
}
