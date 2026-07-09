package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string `env:"DB_HOST" env-default:"localhost"`
	DBPort     int    `env:"DB_PORT" env-default:"5433"`
	DBUser     string `env:"DB_USER" env-default:"postgres"`
	DBPassword string `env:"DB_PASSWORD" env-default:"postgres"`
	DBName     string `env:"DB_NAME" env-default:"user-service"`
	SSLMode    string `env:"DB_SSLMODE" env-default:"disable"`
	Env        string `env:"ENV" env-default:"local"`
	GRPCPort   int    `env:"GRPC_PORT" env-default:"50052"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("⚠️ .env файл не найден, используем системные переменные")
	}

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурации: %w", err)
	}

	return &cfg, nil
}
