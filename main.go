package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
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

	var systemWg sync.WaitGroup

	// For now, just start reading sensors on boot.
	srv.Start()

	hndlr := handler.NewAirQualityHandler(config.GetConfig(), srv)

	handlerServer := hndlr.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	readSensorQuitChan := make(chan bool)

	go srv.RunReadSensors(readSensorQuitChan, &systemWg)

	<-stop

	readSensorQuitChan <- true

	logrus.Info("Stopping program now...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := handlerServer.Shutdown(ctx); err != nil {
		logrus.WithError(err).Error("Shutdown required timeout context")
	}

	systemWg.Wait()

	logrus.Info("Program stopped now")
}
