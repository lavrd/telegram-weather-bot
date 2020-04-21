package language

import (
	"twb/pkg/emoji"
	"twb/pkg/message"
)

const (
	EN = "en"
	RU = "ru"
)

// ru / en language dictionary
var Dictionary = map[string]map[string]string{
	"en": {
		"mps": "mps",
		"mph": "mph",
		"IFL": "feels like",
		"ZERO_RESULTS_LOCATION": "It is impossible to give a " +
			"forecast for the specified coordinates",
		message.Now:              "Now",
		message.ForToday:         "For today",
		message.ForTomorrow:      "For tomorrow",
		message.ForWeek:          "For week",
		message.ChangeLanguageTo: "Change language to",
		"changeUnits":            "Change units to",
		"ZERO_RESULTS_CITY":      "Unable to find the weather for this city",
		"changeCityTo":           "Change city to",
		message.EmptyCity:        "Enter your city to get the actual weather",
		"Monday":                 "Mon",
		"Tuesday":                "Tue",
		"Wednesday":              "Wed",
		"Thursday":               "Thu",
		"Friday":                 "Fri",
		"Saturday":               "Sat",
		"Sunday":                 "Sun",
		"YourLLU":                "Your location, language and units",
		"empty_location":         "Empty location",
		"°C, mps":                "°C, mps",
		"°F, mph":                "°F, mph",
		"mainMenu":               "Main menu",
		"unknownError":           "Unknown error, please try later.",

		message.Help: "Bot show the *current* weather as well as *for today*, *tomorrow* and *the week*.\n\n" +
			"*Now* - current weather\n" +
			"*For today* - weather for today\n" +
			"*For tomorrow* - weather for tomorrow\n" +
			"*For week* - weather for the week\n\n" +
			"You can also send a message to bot or use a bot *commands*.\n\n" +
			"To select the language or units, click " + emoji.Gear + ".\n" +
			"You can select the language by command /lang or by " + emoji.Globe + ".\n" +
			"Also you can select the units by command /units or by " + emoji.TriangularRuler + ".\n" +
			"View the current location, language and units is possible by command /info or " + emoji.Info + ".",
	},

	"ru": {
		"mps":                    "м/с",
		"mph":                    "миль/ч",
		"IFL":                    "Чувствуется как",
		"ZERO_RESULTS_LOCATION":  "Невозможно выдать погоду по данным координатам",
		message.Now:              "Сейчас",
		message.ForToday:         "На сегодня",
		message.ForTomorrow:      "На завтра",
		message.ForWeek:          "На неделю",
		message.ChangeLanguageTo: "Язык изменен на",
		"changeUnits":            "Единицы измерения изменены на",
		"ZERO_RESULTS_CITY":      "Невозможно выдать погоду в данном городе",
		"changeCityTo":           "Город изменен на",
		message.EmptyCity:        "Введите город, чтобы получить актуальную погоду",
		"Monday":                 "Пн",
		"Tuesday":                "Вт",
		"Wednesday":              "Ср",
		"Thursday":               "Чт",
		"Friday":                 "Пт",
		"Saturday":               "Сб",
		"Sunday":                 "Вс",
		"YourLLU":                "Ваша локация, язык и единицы измерения",
		"empty_location":         "Пустое местоположение",
		"°C, mps":                "°C, м/c",
		"°F, mph":                "°F, миль/ч",
		"mainMenu":               "Главное меню",
		"unknownError":           "Неизвестная ошибка, попробуйте позже.",

		message.Help: "Бот показывает *текущую* погоду, а также на *сегодня*, *завтра* и *неделю*.\n\n" +
			"*Сейчас* - текущая погода\n" +
			"*На сегодня* - погода на сегодня\n" +
			"*На завтра* - погода на завтра\n" +
			"*На неделю* - погода на неделю\n\n" +
			"Также вы можете отправить боту сообщение или использовать *команды*.\n\n" +
			"Для выбора языка или единиц измерения, нажмите " + emoji.Gear + ".\n" +
			"Вы можете выбрать язык введя команду /lang или " + emoji.Globe + ".\n" +
			"Также вы можете выбрать единицы измерения введя команду /units или " + emoji.TriangularRuler + ".\n" +
			"Посмотреть текущие настройки локации, языка и единиц измерения можно введя команду /info или " + emoji.Info + ".",
	},
}
