package config

import "os"

type AppConfigStruct struct {
	AppName   string
	AppEnv    string
	AppURL    string
	AppSecret string
	AppPort   string
}

var AppConfig *AppConfigStruct

func LoadAppConfig() {
	AppConfig = &AppConfigStruct{
		AppName:   os.Getenv("APP_NAME"),
		AppEnv:    os.Getenv("APP_ENV"),
		AppURL:    os.Getenv("APP_URL"),
		AppSecret: os.Getenv("APP_SECRET"),
		AppPort:   os.Getenv("APP_PORT"),
	}
}
