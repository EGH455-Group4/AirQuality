package service

import (
	"time"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/models"
	"github.com/ImTheTom/air-quality/sensors"
)

type AirQualityService struct {
	cfg              *config.Config
	Sensors          *models.Sensors
	Errors           []string
	Running          bool
	AirQualityReader *sensors.AirQualityReader
}

func NewAirQualityService(cfg *config.Config, airReader *sensors.AirQualityReader) *AirQualityService {
	return &AirQualityService{
		cfg:              cfg,
		Sensors:          &models.Sensors{},
		Errors:           []string{},
		Running:          false,
		AirQualityReader: airReader,
	}
}

func (s *AirQualityService) GetAirQuality() *models.AirQuality {
	return &models.AirQuality{
		Sensors:     s.Sensors,
		CurrentTime: time.Now(),
		Errors:      s.Errors,
	}
}

func (s *AirQualityService) SingleRead() *models.AirQuality {
	s.ReadSensors()

	return &models.AirQuality{
		Sensors:     s.Sensors,
		CurrentTime: time.Now(),
		Errors:      s.Errors,
	}
}

func (s *AirQualityService) Start() {
	s.ResetVars()
	s.Running = true
}

func (s *AirQualityService) Stop() {
	s.ResetVars()
	s.Running = false
}

func (s *AirQualityService) ResetVars() {
	s.Sensors = &models.Sensors{}
	s.Errors = []string{}
}

func (s *AirQualityService) ReadSensors() {
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