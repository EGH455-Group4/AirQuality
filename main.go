package main

import (
	"github.com/ImTheTom/air-quality/config"
	"github.com/sirupsen/logrus"
)

func main() {
	config.LoggerSetup()

	config.SetupConfig("config.json")

	logrus.WithField("config", config.GetConfig()).Info("Running with config")

	logrus.Info("Hello world")
}
