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

	cfg := config.GetConfig()

	logrus.WithField("config", cfg).Info("Running with config")

	localIP := getLocalIP()

	logrus.WithField("local_ip", localIP).Infof("running on: %s", localIP+cfg.Address)

	reader := sensors.NewAirQualityReader(cfg)

	srv := service.NewAirQualityService(cfg, reader)

	var systemWg sync.WaitGroup

	if cfg.StartOnBoot {
		srv.Start()
	}

	hndlr := handler.NewAirQualityHandler(cfg, srv)

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
