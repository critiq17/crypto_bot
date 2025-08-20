package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("File .env not found")
	}

	return &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
	}
}
