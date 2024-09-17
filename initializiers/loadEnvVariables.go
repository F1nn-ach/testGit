package initializiers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load("initializiers/port.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
