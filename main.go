package main

import (
	"context"
	"os"
	"os/signal"
	"time"

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

	handlerServer := hndlr.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := handlerServer.Shutdown(ctx); err != nil {
		logrus.WithError(err).Error("Shutdown required timeout context")
	}
}
