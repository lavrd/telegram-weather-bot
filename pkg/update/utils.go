package update

import (
	"strings"

	"twb/pkg/emoji"
	"twb/pkg/message"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
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

func parseIncomingMsg(msg, cmd string) MsgType {
	text := ""
	if msg != "" {
		text = msg
	} else {
		text = cmd
	}
	if strings.Contains(text, "/") {
		text = text[1:]
	}

	switch text {
	case message.Start:
		return StartMsg
	case emoji.Back:
		return BackMsg
	case emoji.CountriesFATE[language.English.String()], emoji.CountriesFATE[language.Russian.String()]:
		return UpdateLangMsg
	case emoji.Globe, message.Lang:
		return langKeyboardMsg
	case emoji.Help, message.Help:
		return HelpMsg
	default:
		return UnknownMsg
	}
}
