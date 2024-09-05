package pkg

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load(`.env`)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
