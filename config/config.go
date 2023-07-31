package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

type Config struct {
	MockHardware      bool   `json:"mock_hardware"`
	Address           string `json:"address"`
	SensorReadSeconds int    `json:"sensor_read_seconds"`
	ParallelRead      bool   `json:"parallel_read"`
	StartOnBoot       bool   `json:"start_on_boot"`
}

var (
	cfg *Config

	errNoConfig = errors.New("Can't find config file")
)

func GetConfig() *Config {
	return cfg
}

func SetupConfig(location string) {
	defaultConfig := Config{
		MockHardware:      false,
		Address:           ":8050",
		SensorReadSeconds: 5,
		ParallelRead:      false,
		StartOnBoot:       false,
	}

	cfg = &defaultConfig

	if len(location) == 0 {
		logrus.Warn(errNoConfig)

		return
	}

	raw, err := ioutil.ReadFile(location)
	if err != nil {
		logrus.Warn(err)

		return
	}

	if err = json.Unmarshal(raw, &defaultConfig); err != nil {
		logrus.Warn(err)
	}
}
