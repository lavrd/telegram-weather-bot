package update

import (
	"fmt"
	"math"

	"twb/pkg/emoji"
	"twb/pkg/forecast"
	"twb/pkg/language"
	"twb/pkg/message"
	"twb/pkg/storage"
)

func getCurrentWeather(forecast *forecast.Current, user *storage.User) string {
	return fmt.Sprintf(
		"%s *%f%s %s* `%s\n%s` *%f%s*",
	)
	// return model.Icons[f.Currently.Icon] + " *" + format.FTS0(f.Currently.Temperature) +
	// 	getTempUnit(units) + "  " +
	// 	getWind(f.Currently.WindSpeed, f.Currently.WindBearing, lang, units) +
	// 	"*  `" + f.Currently.Summary + ".`\n`" + language.Language[lang]["IFL"] +
	// 	"`  *" + format.FTS0(f.Currently.ApparentTemperature) + getTempUnit(units) + "*"
}

func prepareConditions(conditions forecast.Conditions) {
	if len(conditions) == 1 {

	}
	for _, condition := range conditions {
		_ = condition
	}
}

func prepareTempUnits(units storage.Units) string {
	if units == storage.SI {
		return message.C
	}
	return message.F
}

func prepareWind(wind *forecast.Wind, units storage.Units, lang string) string {
	sector := int(math.Mod(360+float64(wind.Deg)/22.5+.5, 16))
	return fmt.Sprintf("%s %s %s", emoji.Directions[sector], wind.Speed, prepareWindUnit(units, lang))
}

func prepareWindUnit(units storage.Units, lang string) string {
	if units == storage.SI {
		return language.Dictionary[lang][message.MPS]
	}
	return language.Dictionary[lang][message.MPH]
}
