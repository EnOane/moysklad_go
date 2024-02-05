package config

import (
	"os"
	"strconv"
)

type PostgresConfig struct {
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
}

type MoyskladConfig struct {
	Username string
	Password string
}

type Config struct {
	MoyskladConfig MoyskladConfig
	PostgresConfig PostgresConfig
}

var AppConfig Config

func NewConfig() {
	AppConfig = Config{
		MoyskladConfig: MoyskladConfig{
			Username: getEnv("MOYSKLAD_USERNAME", ""),
			Password: getEnv("MOYSKLAD_PASSWORD", ""),
		},
		PostgresConfig: PostgresConfig{
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnvAsInt("DB_PORT", 5432),
			DbName:   getEnv("DB_NAME", ""),
			User:     getEnv("DB_USER", ""),
			Password: getEnv("DB_PASSWORD", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
