package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Url string
}

func (c *Config) Init(configFile string) {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Println("Error loading from .env file, procceed...")
	}
	c.Url = os.Getenv("SERVER_URL")
}
