package update

import (
	"twb/pkg/emoji"
	twbl "twb/pkg/language"
	"twb/pkg/message"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/text/language"
)

func mainKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(twbl.Dictionary[lang][message.Now]),
			tgbotapi.NewKeyboardButton(twbl.Dictionary[lang][message.ForToday]),
		},
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(twbl.Dictionary[lang][message.ForTomorrow]),
			tgbotapi.NewKeyboardButton(twbl.Dictionary[lang][message.ForWeek]),
		},
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(emoji.Gear),
			tgbotapi.NewKeyboardButton(emoji.Info),
			tgbotapi.NewKeyboardButton(emoji.Help),
		},
	)
}

// func unitsKeyboard(lang string) tgbotapi.ReplyKeyboardMarkup {
// 	return tgbotapi.NewReplyKeyboard(
// 		[]tgbotapi.KeyboardButton{
// 			tgbotapi.NewKeyboardButton(model.Back),
// 			tgbotapi.NewKeyboardButton(twbl.Languages[lang]["°C, mps"]),
// 			tgbotapi.NewKeyboardButton(twbl.Languages[lang]["°F, mph"]),
// 		},
// 	)
// }

// func settingsKeyboard() tgbotapi.ReplyKeyboardMarkup {
// 	return tgbotapi.NewReplyKeyboard(
// 		[]tgbotapi.KeyboardButton{
// 			tgbotapi.NewKeyboardButton(model.Back),
// 			tgbotapi.NewKeyboardButton(model.GlobeWithMeridian),
// 			tgbotapi.NewKeyboardButton(model.TriangularRuler),
// 		},
// 	)
// }

func langKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(emoji.Back),
			tgbotapi.NewKeyboardButton(emoji.CountriesFATE[language.English.String()]),
			tgbotapi.NewKeyboardButton(emoji.CountriesFATE[language.Russian.String()]),
		},
	)
}

func helpKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		[]tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(emoji.Help),
		},
	)
}
