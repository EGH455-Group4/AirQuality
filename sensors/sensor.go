package sensors

import (
	"errors"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/models"
)

type AirQualityReader struct {
	cfg              *config.Config
	AirQualitySensor AirQualitySensor
}

var ErrorUnknownReading = errors.New("Unknown wanted reading")

func NewAirQualityReader(cfg *config.Config) *AirQualityReader {
	airQualitySensor := NewEnviroSensor()
	if cfg.MockHardware {
		airQualitySensor = NewFakeSensor()
	}

	return &AirQualityReader{
		cfg:              cfg,
		AirQualitySensor: airQualitySensor,
	}
}

func (s *AirQualityReader) ReadSensor(wantedReading Reading) (*models.SensorReading, error) {
	var (
		reading float64
		err     error
	)

	switch wantedReading {
	case Light:
		reading, err = s.AirQualitySensor.ReadLight()
	case HazardousGases:
		reading, err = s.AirQualitySensor.ReadHazardousGases()
	case Humidity:
		reading, err = s.AirQualitySensor.ReadHumidity()
	case Pressure:
		reading, err = s.AirQualitySensor.ReadPressure()
	case Temperature:
		reading, err = s.AirQualitySensor.ReadTemperature()
	default:
		return nil, ErrorUnknownReading
	}

	if err != nil {
		return nil, err
	}

	return &models.SensorReading{
		Reading: reading,
	}, nil
}
