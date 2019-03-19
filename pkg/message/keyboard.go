package msg

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	l "telegram-weather-bot/pkg/language"
	"telegram-weather-bot/pkg/model"
	"golang.org/x/text/language"
)

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

func unitsKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(model.Back),
			tgbotapi.NewKeyboardButton(l.Language[lang]["°C, mps"]),
			tgbotapi.NewKeyboardButton(l.Language[lang]["°F, mph"]),
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
