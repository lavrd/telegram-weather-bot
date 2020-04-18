package geocode

type Result struct {
	Addr     string
	Lat, Lon float64
}

type Geocode interface {
	Geocode(location, lang string) (*Result, error)
	Reverse(lat, lon float64, lang string) (*Result, error)
}
