package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Driver   string
}

type ApiConfig struct {
	ApiPort string
}
type Config struct {
	DbConfig
	ApiConfig
}

func (c *Config) ConfigConfiguration() error {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	//config db
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}
	//config Port apps

	if c.Host == "" || c.Port == "" {
		return fmt.Errorf("missing requirement")
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.ConfigConfiguration(); err != nil {
		panic(err)
	}

	return cfg, nil
}
