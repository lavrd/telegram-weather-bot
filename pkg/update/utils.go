package update

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	LoggerFieldKeyMessage    = "message"
	LoggerFieldKeyTelegramID = "telegramId"
)

func prepareLogger(telegramID int64, msg string) zerolog.Logger {
	return log.Logger.
		With().
		Int64(LoggerFieldKeyTelegramID, telegramID).
		Str(LoggerFieldKeyMessage, msg).
		Logger()
}
