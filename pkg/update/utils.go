package update

import (
	"strings"

	"twb/pkg/emoji"
	twbl "twb/pkg/language"
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
	case
		strings.ToLower(twbl.Dictionary[twbl.EN][message.Now]), strings.ToLower(twbl.Dictionary[twbl.RU][message.Now]),
		strings.ToLower(twbl.Dictionary[twbl.EN][message.ForToday]), strings.ToLower(twbl.Dictionary[twbl.RU][message.ForToday]),
		strings.ToLower(twbl.Dictionary[twbl.EN][message.ForTomorrow]), strings.ToLower(twbl.Dictionary[twbl.RU][message.ForTomorrow]),
		strings.ToLower(twbl.Dictionary[twbl.EN][message.ForWeek]), strings.ToLower(twbl.Dictionary[twbl.RU][message.ForWeek]):
		return WeatherFromCmd

	case message.Start:
		return StartMsg

	case emoji.Back:
		return BackMsg

	case emoji.CountriesFATE[language.English.String()], emoji.CountriesFATE[language.Russian.String()]:
		return LangUpdateMsg

	case emoji.Globe, message.Lang:
		return langKeyboardMsg

	case emoji.Help, message.Help:
		return HelpMsg

	default:
		return AnyMsg
	}
}
