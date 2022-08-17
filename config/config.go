package config

import (
	"fmt"
	"os"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	DataSourceName string
}

type GrpcConfig struct {
	Url string
}

type Config struct {
	DbConfig
	ApiConfig
}

func (c *Config) readConfig() {
	// API Config start here
	api := os.Getenv("API_URL")
	c.ApiConfig = ApiConfig{Url: api}

	// DB Config start here
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)
	c.DbConfig = DbConfig{dsn}

}

func InitConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
