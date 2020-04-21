package emoji

var (
	// emoji -> lang
	CountriesFETA = map[string]string{
		"🇷🇺": "ru",
		"🇬🇧": "en",
	}

	// lang -> emoji
	CountriesFATE = map[string]string{
		"ru": "🇷🇺",
		"en": "🇬🇧",
	}

	Help            = "🆘"
	Globe           = "🌏"
	Info            = "ℹ️"
	Sunrise         = "🌄"
	Sunset          = "🌆"
	Gear            = "⚙️"
	TriangularRuler = "📐"
	Back            = "↩️"

	Moons = map[string]string{
		"new moon":           "🌑",
		"first quarter moon": "🌓",
		"full moon":          "🌕",
		"last quarter moon":  "🌗",
	}

	Icons = map[string]string{
		"clear-day":           "\xE2\x98\x80",
		"clear-night":         "\xF0\x9F\x8C\x99",
		"partly-cloudy-day":   "\xE2\x9B\x85",
		"partly-cloudy-night": "\xE2\x98\x81",
		"cloudy":              "\xE2\x98\x81",
		"rain":                "\xF0\x9F\x8C\xA7",
		"sleet":               "\xF0\x9F\x8C\xA7",
		"snow":                "\xF0\x9F\x8C\xA8",
		"wind":                "\xF0\x9F\x92\xA8",
		"fog":                 "\xF0\x9F\x8C\xAB",
	}

	Directions = [16]string{
		"\xE2\xAC\x86",
		"\xE2\x86\x97",
		"\xE2\x86\x97",
		"\xE2\x86\x97",
		"\xE2\x9E\xA1",
		"\xE2\x86\x98",
		"\xE2\x86\x98",
		"\xE2\x86\x98",
		"\xE2\xAC\x87",
		"\xE2\x86\x99",
		"\xE2\x86\x99",
		"\xE2\x86\x99",
		"\xE2\xAC\x85",
		"\xE2\x86\x96",
		"\xE2\x86\x96",
		"\xE2\x86\x96",
	}
)
