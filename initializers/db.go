package initializers

import (
	"fmt"
	"go-auth/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToPgSql() *gorm.DB {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DatabaseConfig.DBHost,
		config.DatabaseConfig.DBUser,
		config.DatabaseConfig.DBPassword,
		config.DatabaseConfig.DBName,
		config.DatabaseConfig.DBPort,
		config.DatabaseConfig.SSLMode,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return DB
}
