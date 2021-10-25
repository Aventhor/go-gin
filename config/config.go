package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT     string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

func New() *Config {
	return &Config{PORT: getEnv("PORT"), Database: DatabaseConfig{Host: getEnv("DB_HOST"), Name: getEnv("POSTGRES_DB"), User: getEnv("POSTGRES_USER"), Password: getEnv("POSTGRES_PASSWORD")}}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if exists == false {
		fmt.Println(key, "missing in .env")
		return ""
	}

	return value
}

func Load() {
	err := godotenv.Load("./.env")

	if err != nil {
		panic(err)
	}
}
