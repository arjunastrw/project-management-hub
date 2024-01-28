package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
type TokenConfig struct {
	Issuer           string `json:"issuer"`
	JwtSignatureKey  []byte `json:"jwt_signature_key"`
	JwtSigningMethod *jwt.SigningMethodHMAC
	JwtExpiresTime   time.Duration
}

type PathConfig struct {
	StaticPath string `json:"static_path"`
}

type Config struct {
	DbConfig
	ApiConfig
	TokenConfig
	PathConfig
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

	tokenExpire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE"))
	c.TokenConfig = TokenConfig{
		Issuer:           os.Getenv("JWT_ISSUER"),
		JwtSignatureKey:  []byte(os.Getenv("JWT_SIGNATURE_KEY")),
		JwtSigningMethod: jwt.SigningMethodHS256,
		JwtExpiresTime:   time.Duration(tokenExpire) * time.Hour,
	}

	c.PathConfig = PathConfig{StaticPath: os.Getenv("FILE_PATH")}
	if c.PathConfig.StaticPath == "" {
		return fmt.Errorf("missing requirement FILE_PATH in .env ")
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
