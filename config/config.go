package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	LogLevel   string `env:"LOG_LEVEL,required"`
	Mode       string `env:"GIN_MODE",required"`
	ServerPort string `env:"SERVER_PORT"`
	ServerHost string `env:"SERVER_HOST,required"`

	DBHost     string `env:"DB_HOST,required"`
	DBPort     string `env:"DB_PORT,required"`
	DBName     string `env:"DB_NAME,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBSSLMode  string `env:"DB_SSL"`

	TokenAccessTime  int    `env:"TOKEN_ACCESS_TIME,required"`
	TokenRefreshTime int    `env:"TOKEN_REFRESH_TIME,required"`
	TokenVerifyTime  int    `env:"TOKEN_VERIFY_TIME,required"`
	TokenSecret      string `env:"TOKEN_SECRET,required"`

	SmtpServer     string `env:"SMTP_SERVER,required"`
	SmtpPort       int    `env:"SMTP_PORT,required"`
	SmtpUser       string `env:"SMTP_USER,required"`
	SmtpPassword   string `env:"SMTP_PASSWORD,required"`
	SenderIdentity string `env:"SENDER_IDENTITY,required"`
	SenderEmail    string `env:"SENDER_EMAIL,required"`
}

func NewEnvConfig() *EnvConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Unable to load .env: %e", err)
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to load variables from .env: %e", err)
	}

	return config
}
