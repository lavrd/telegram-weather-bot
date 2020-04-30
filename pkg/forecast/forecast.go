package forecast

type Forecast interface {
	GetCurrent(lat, lon float64, units, lang string) (*Current, error)
}

type Current struct {
	Temp         string
	ApparentTemp string
	Conditions   Conditions
	Wind         *Wind
}

type Wind struct {
	Speed string
	Deg   int
}

type Condition struct {
	Summary string
	Type    ConditionType
}
type Conditions []*Condition

type ConditionType int

const (
	Unknown ConditionType = iota
	Thunderstorm
	Drizzle
	Rain
	Snow
	Mist
	Smoke
	Haze
	Dust
	Fog
	Sand
	Ash
	Squall
	Tornado
	Clear
	Clouds
)
