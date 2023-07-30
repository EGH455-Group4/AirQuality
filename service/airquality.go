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
	Errors           []string
	Running          bool
	AirQualityReader sensors.AirQualityReader
}

func NewAirQualityService(cfg *config.Config, airReader sensors.AirQualityReader) AirQualityService {
	return &airQualityService{
		cfg:              cfg,
		Sensors:          &models.Sensors{},
		Errors:           []string{},
		Running:          false,
		AirQualityReader: airReader,
	}
}

func (s *airQualityService) GetAirQuality() *models.AirQuality {
	return &models.AirQuality{
		Sensors:     s.Sensors,
		CurrentTime: time.Now(),
		Errors:      s.Errors,
	}
}

func (s *airQualityService) SingleRead() *models.AirQuality {
	s.readSensors()

	return &models.AirQuality{
		Sensors:     s.Sensors,
		CurrentTime: time.Now(),
		Errors:      s.Errors,
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
	logrus.Info("Running read sensors loop")

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

				logrus.WithField("sensor_reading", s.Sensors).Info("Read sensors")
			} else {
				logrus.Info("Not currently running read sensors")
			}

			time.Sleep(time.Duration(s.cfg.SensorReadSeconds) * time.Second)
		}
	}
}

func (s *airQualityService) readSensors() {
	s.Errors = []string{}

	if lightReading, err := s.AirQualityReader.ReadSensor(sensors.Light); err != nil {
		s.Errors = append(s.Errors, err.Error())
		s.Sensors.Light = nil
	} else {
		s.Sensors.Light = lightReading
	}

	if hazardousGasesReading, err := s.AirQualityReader.ReadSensor(sensors.HazardousGases); err != nil {
		s.Errors = append(s.Errors, err.Error())
		s.Sensors.HazardousGases = nil
	} else {
		s.Sensors.HazardousGases = hazardousGasesReading
	}

	if humidityReading, err := s.AirQualityReader.ReadSensor(sensors.Humidity); err != nil {
		s.Errors = append(s.Errors, err.Error())
		s.Sensors.Humidity = nil
	} else {
		s.Sensors.Humidity = humidityReading
	}

	if pressureReading, err := s.AirQualityReader.ReadSensor(sensors.Pressure); err != nil {
		s.Errors = append(s.Errors, err.Error())
		s.Sensors.Pressure = nil
	} else {
		s.Sensors.Pressure = pressureReading
	}

	if temperatureReading, err := s.AirQualityReader.ReadSensor(sensors.Temperature); err != nil {
		s.Errors = append(s.Errors, err.Error())
		s.Sensors.Temperature = nil
	} else {
		s.Sensors.Temperature = temperatureReading
	}
}

func (s *airQualityService) resetVars() {
	s.Sensors = &models.Sensors{}
	s.Errors = []string{}
}
