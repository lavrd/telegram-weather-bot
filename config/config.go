package config

import (
	"github.com/go-yaml/yaml"
	m "github.com/lavrs/telegram-weather-bot/model"
	"io/ioutil"
	"log"
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

	if file, err = ioutil.ReadFile("config.yml"); err != nil {
		log.Panic(err)
	}

	if err := yaml.Unmarshal(file, &data); err != nil {
		log.Panic(err)
	}

	return &data
}
