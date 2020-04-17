package config

var Cfg = &struct {
	Google struct {
		Geocoding struct {
			Token string
		}
	}
	// OpenWeatherMap config
	OWM struct {
		Token string
	}
	Telegram struct {
		Token string
		Error struct {
			Admin int64
		}
	}
	DSN string
}{}
