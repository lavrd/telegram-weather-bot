package model

// bot config
type Config struct {
	TelegramToken        string `yaml:"telegram-token"`
	TelegramTestToken    string `yaml:"telegram-test-token"`
	DarkskyToken         string `yaml:"darksky-token"`
	GoogleGeocodingToken string `yaml:"google-geocoding-token"`
	MyTelegramID         int64  `yaml:"my-telegram-id"`
}

// db struct
type DB struct {
	TelegramID int64   `gorethink:"telegramID"`
	Location   string  `gorethink:"location"`
	Lang       string  `gorethink:"lang"`
	Lat        float64 `gorethink:"lat"`
	Lng        float64 `gorethink:"lng"`
	Units      string  `gorethink:"units"`
}

var (
	// emoji -> lang
	CountriesFETA = map[string]string{
		"\xF0\x9F\x87\xB7\xF0\x9F\x87\xBA": "ru",
		"\xF0\x9F\x87\xAC\xF0\x9F\x87\xA7": "en",
	}

	// lang -> emoji
	CountriesFATE = map[string]string{
		"ru": "\xF0\x9F\x87\xB7\xF0\x9F\x87\xBA",
		"en": "\xF0\x9F\x87\xAC\xF0\x9F\x87\xA7",
	}

	// emoji
	Help              = "\xF0\x9F\x86\x98"
	GlobeWithMeridian = "\xF0\x9F\x8C\x90"
	Info              = "\xE2\x84\xB9"
	Sunrise           = "\xF0\x9F\x8C\x85"
	Sunset            = "\xF0\x9F\x8C\x84"
	Gear              = "\xE2\x9A\x99"
	TriangularRuler   = "\xF0\x9F\x93\x90"
	Back              = "\xE2\x97\x80"

	// moons emoji
	Moons = map[string]string{
		"new moon":           "\xF0\x9F\x8C\x91",
		"first quarter moon": "\xF0\x9F\x8C\x93",
		"full moon":          "\xF0\x9F\x8C\x95",
		"last quarter moon":  "\xF0\x9F\x8C\x97",
	}

	// weather icons emoji
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

	// directions emoji
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
