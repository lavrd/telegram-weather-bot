package config

var (
	Viper = &struct {
		Google struct {
			Geocoding struct {
				Token string
			}
		}
		Darksky struct {
			Token string
		}
		Telegram struct {
			Error struct {
				Admin int64
				Send  bool
			}
			Token string
		}
		Database struct {
			Endpoint string
		}
	}{}
)
