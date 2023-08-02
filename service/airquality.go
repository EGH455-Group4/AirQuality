package service

import (
	"sync"
	"time"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/models"
	"github.com/ImTheTom/air-quality/sensors"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination mocks/airquality.mock.go -package mocks -source airquality.go

type AirQualityService interface {
	GetAirQuality() *models.AirQuality
	SingleRead() *models.AirQuality
	Start()
	Stop()
	RunReadSensors(closeCh chan bool, wg *sync.WaitGroup)
}

type airQualityService struct {
	cfg              *config.Config
	Sensors          *models.Sensors
	Running          bool
	AirQualityReader sensors.AirQualityReader
	ReadWg           *sync.WaitGroup
	ReadTime         time.Time
}

func NewAirQualityService(cfg *config.Config, airReader sensors.AirQualityReader) AirQualityService {
	ReadWg := &sync.WaitGroup{}

	return &airQualityService{
		cfg:              cfg,
		Sensors:          &models.Sensors{},
		Running:          false,
		AirQualityReader: airReader,
		ReadWg:           ReadWg,
		ReadTime:         time.Time{},
	}
}

func (s *airQualityService) GetAirQuality() *models.AirQuality {
	return &models.AirQuality{
		Sensors:     s.Sensors,
		CurrentTime: s.ReadTime,
	}
}

func (s *airQualityService) SingleRead() *models.AirQuality {
	s.readSensors()

	return &models.AirQuality{
		Sensors:     s.Sensors,
		CurrentTime: s.ReadTime,
	}
}

func (s *airQualityService) Start() {
	s.resetVars()
	s.Running = true
}

func (s *airQualityService) Stop() {
	s.resetVars()
	s.Running = false
}

func (s *airQualityService) RunReadSensors(closeCh chan bool, wg *sync.WaitGroup) {
	logrus.WithField("parallel", s.cfg.ParallelRead).Info("Running read sensors loop")

	wg.Add(1)
	defer wg.Done()

	for {
		select {
		case <-closeCh:
			close(closeCh)

			logrus.Info("Received shutdown call, exiting now...")

			return
		default:
			if s.Running {
				s.readSensors()

				logrus.WithFields(logrus.Fields{
					"sensor_reading": s.Sensors,
				}).Info("Read sensors")
			}

			time.Sleep(time.Duration(s.cfg.SensorReadSeconds) * time.Second)
		}
	}
}

func (s *airQualityService) readSensors() {
	if s.cfg.ParallelRead {
		s.readSensorsParallel()

		s.ReadTime = time.Now()

		return
	}

	s.readSensorsLinear()

	s.ReadTime = time.Now()
}

func (s *airQualityService) readSensorsLinear() {
	s.Sensors.Light = s.AirQualityReader.ReadSensor(sensors.Light)

	s.Sensors.HazardousGases = s.AirQualityReader.ReadSensor(sensors.HazardousGases)

	s.Sensors.Humidity = s.AirQualityReader.ReadSensor(sensors.Humidity)

	s.Sensors.Pressure = s.AirQualityReader.ReadSensor(sensors.Pressure)

	s.Sensors.Temperature = s.AirQualityReader.ReadSensor(sensors.Temperature)
}

func (s *airQualityService) readSensorsParallel() {
	s.ReadWg.Add(5)

	go func(syncWg *sync.WaitGroup) {
		defer syncWg.Done()

		s.Sensors.Light = s.AirQualityReader.ReadSensor(sensors.Light)
	}(s.ReadWg)

	go func(syncWg *sync.WaitGroup) {
		defer syncWg.Done()

		s.Sensors.HazardousGases = s.AirQualityReader.ReadSensor(sensors.HazardousGases)
	}(s.ReadWg)

	go func(syncWg *sync.WaitGroup) {
		defer syncWg.Done()

		s.Sensors.Humidity = s.AirQualityReader.ReadSensor(sensors.Humidity)
	}(s.ReadWg)

	go func(syncWg *sync.WaitGroup) {
		defer syncWg.Done()

		s.Sensors.Pressure = s.AirQualityReader.ReadSensor(sensors.Pressure)
	}(s.ReadWg)

	go func(syncWg *sync.WaitGroup) {
		defer syncWg.Done()

		s.Sensors.Temperature = s.AirQualityReader.ReadSensor(sensors.Temperature)
	}(s.ReadWg)

	s.ReadWg.Wait()
}

func (s *airQualityService) resetVars() {
	s.Sensors = &models.Sensors{}
}
