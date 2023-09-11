package config

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvVariableInitilizer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}
}
