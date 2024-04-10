package utils

import (
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

// LoadEnv attempts to load environment variables from a .env file located at the root of the project.
// This allows for easier management of environment-specific settings across different deployment environments.
// Utilizes the godotenv library to parse the .env file and load variables into the Go process's environment.
// If the .env file cannot be found or read, the function logs a panic message and the application will likely terminate.
// This behavior ensures that the application does not run with incomplete or incorrect configuration.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Panic().Msg("---failed to load .env file---")
	}
}
