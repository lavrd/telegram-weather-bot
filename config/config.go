package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	m "github.com/spacelavr/telegram-weather-bot/model"
)

// Cfg store config
var Cfg m.Config

// SetConfig setting bot config
func SetConfig() {
	data := openFile()
	&Cfg = data
}

// open config file
func openFile() *m.Config {
	var (
		file []byte
		err  error
		data m.Config
	)

	if file, err = ioutil.ReadFile("config.json"); err != nil {
		log.Panic(err)
	}

	if err := json.Unmarshal(file, &data); err != nil {
		log.Panic(err)
	}

	return &data
}
