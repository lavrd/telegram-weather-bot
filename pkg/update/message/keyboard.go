package message

import (
	twbl "telegram-weather-bot/pkg/language"
	"telegram-weather-bot/pkg/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/text/language"
)

func mainKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(twbl.Languages[lang]["now"]),
			tgbotapi.NewKeyboardButton(twbl.Languages[lang]["forToday"]),
		},
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(twbl.Languages[lang]["forTomorrow"]),
			tgbotapi.NewKeyboardButton(twbl.Languages[lang]["forWeek"]),
		},
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Gear),
			tgbotapi.NewKeyboardButton(model.Info),
			tgbotapi.NewKeyboardButton(model.Help),
		},
	)
}

func unitsKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Back),
			tgbotapi.NewKeyboardButton(twbl.Languages[lang]["°C, mps"]),
			tgbotapi.NewKeyboardButton(twbl.Languages[lang]["°F, mph"]),
		},
	)
}

func settingsKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Back),
			tgbotapi.NewKeyboardButton(model.GlobeWithMeridian),
			tgbotapi.NewKeyboardButton(model.TriangularRuler),
		},
	)
}

func langKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Back),
			tgbotapi.NewKeyboardButton(model.CountriesFATE[language.English.String()]),
			tgbotapi.NewKeyboardButton(model.CountriesFATE[language.Russian.String()]),
		},
	)
}

func helpKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Help),
		},
	)
}
