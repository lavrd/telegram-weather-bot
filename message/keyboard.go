package msg

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	l "github.com/spacelavr/telegram-weather-bot/language"
	"github.com/spacelavr/telegram-weather-bot/model"
	"golang.org/x/text/language"
)

// main bot keyboard
func mainKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(l.Language[lang]["now"]),
			tgbotapi.NewKeyboardButton(l.Language[lang]["forToday"]),
		},
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(l.Language[lang]["forTomorrow"]),
			tgbotapi.NewKeyboardButton(l.Language[lang]["forWeek"]),
		},
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Gear),
			tgbotapi.NewKeyboardButton(model.Info),
			tgbotapi.NewKeyboardButton(model.Help),
		},
	)
}

// units keyboard
func unitsKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Back),
			tgbotapi.NewKeyboardButton(l.Language[lang]["°C, mps"]),
			tgbotapi.NewKeyboardButton(l.Language[lang]["°F, mph"]),
		},
	)
}

// settings keyboard
func settingsKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Back),
			tgbotapi.NewKeyboardButton(model.GlobeWithMeridian),
			tgbotapi.NewKeyboardButton(model.TriangularRuler),
		},
	)
}

// language keyboard
func langKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Back),
			tgbotapi.NewKeyboardButton(model.CountriesFATE[language.English.String()]),
			tgbotapi.NewKeyboardButton(model.CountriesFATE[language.Russian.String()]),
		},
	)
}

// help keyboard
func helpKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Help),
		},
	)
}
