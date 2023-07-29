package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	MockHardware bool `json:"mock_hardware"`
}

var (
	cfg *Config
)

var (
	errBadConfig = errors.New("Bad config was loaded in")
	errNoConfig  = errors.New("Can't find config file")
)

func GetConfig() *Config {
	return cfg
}

func SetupConfig(location string) {
	defaultConfig := Config{
		MockHardware: false,
	}

	cfg = &defaultConfig

	if len(location) == 0 {
		logrus.Warn(errNoConfig)

		return
	}

	raw, err := os.ReadFile(location)
	if err != nil {
		logrus.Warn(err)

		return
	}

	if err = json.Unmarshal(raw, &defaultConfig); err != nil {
		logrus.Warn(err)
	}
}
