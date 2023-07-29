package main

import (
	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	config.LoggerSetup()

	config.SetupConfig("config.json")

	logrus.WithField("config", config.GetConfig()).Info("Running with config")

	hndlr := handler.NewAirQualityHandler(config.GetConfig())

	if err := hndlr.Run(); err != nil {
		logrus.Fatal(err.Error())
	}
}
