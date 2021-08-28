package server

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once

type Config struct {
	settings *Settings
	database string
}

func NewConfig() *Config {
	config := Config{
		settings: NewSettings(),
		database: "",
	}
	initLogger(config.settings.Debug)

	return &config
}

func (config *Config) Settings() *Settings {
	return config.settings
}

func initLogger(debug bool) {
	once.Do(func() {
		Log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})

		if debug {
			Log.SetLevel(logrus.DebugLevel)
		} else {
			Log.SetLevel(logrus.InfoLevel)
		}
	})
}
