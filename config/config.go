package config

import (
	"os"
)

type ConfigStruct struct {
	ServerPort string
	DbDriver   string
	DbSource   string
}

var Config ConfigStruct

func Load() {
	Config = ConfigStruct{
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		DbDriver:   getEnv("DB_DRIVER", "sqlite"),
		DbSource:   getEnv("DB_SOURCE", "employees.db"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
