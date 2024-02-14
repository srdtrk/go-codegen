package types

import (
	"os"

	"github.com/rs/zerolog"
)

func DefaultLogger() *zerolog.Logger {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr},
	).Level(zerolog.TraceLevel).With().Timestamp().Logger()

	return &logger
}
