package main

import (
	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/handler"
	"github.com/ImTheTom/air-quality/sensors"
	"github.com/ImTheTom/air-quality/service"
	"github.com/sirupsen/logrus"
)

func main() {
	config.LoggerSetup()

	config.SetupConfig("config.json")

	logrus.WithField("config", config.GetConfig()).Info("Running with config")

	reader := sensors.NewAirQualityReader(config.GetConfig())

	srv := service.NewAirQualityService(config.GetConfig(), reader)

	hndlr := handler.NewAirQualityHandler(config.GetConfig(), srv)

	if err := hndlr.Run(); err != nil {
		logrus.Fatal(err.Error())
	}
}
