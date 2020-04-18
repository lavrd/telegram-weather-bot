package openweathermap

type OpenWeatherMap struct {
	token string
}

func (owm *OpenWeatherMap) GetToday()    {}
func (owm *OpenWeatherMap) GetTomorrow() {}

func New(token string) *OpenWeatherMap {
	return &OpenWeatherMap{
		token: token,
	}
}
