package config

import "os"

type DatabaseConfigStruct struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	SSLMode    string
}

var DatabaseConfig *DatabaseConfigStruct

func LoadDatabaseConfig() {
	DatabaseConfig = &DatabaseConfigStruct{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		SSLMode:    os.Getenv("SSL_MODE"),
	}
}
