package utils

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// InitLogger initializes the global logger configuration using the zerolog package.
// It sets the global logging level to InfoLevel, ensuring that only logs of level Info and above are printed.
// It configures the logger to output logs in a console-friendly format to stderr,
// and sets up error stack trace marshaling for enhanced error reporting.
//
// This function is typically called at the application startup to configure logging according to these predefined settings.
// It's crucial for maintaining a consistent logging strategy across the application.
func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
