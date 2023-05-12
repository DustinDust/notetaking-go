package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Url string
	DB  string
}

func (c *Config) Init(configFile string) {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Println("Error loading from .env file, procceed...")
	}
	c.Url = strings.TrimSpace(os.Getenv("SERVER_URL"))
	c.DB = strings.TrimSpace(os.Getenv("DB_URI"))
}
