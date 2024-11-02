package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	Telegram TelegramConfig
}

type TelegramConfig struct {
	APIToken string
}

func LoadConfig() Config {
	return Config{
		Telegram: TelegramConfig{
			APIToken: os.Getenv("TELEGRAM_API_TOKEN"),
		},
	}
}
