package loggers

import (
	"os"

	"github.com/rs/zerolog"
)

// Log to log msg
func Log(msg string) {
	logger := zerolog.New(os.Stderr).With().Caller().Logger()
	logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger.Debug().Msg(msg)
}
