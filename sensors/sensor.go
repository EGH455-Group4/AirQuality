package sensors

import (
	"errors"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/models"
)

type AirQualityReader interface {
	ReadSensor(wantedReading Reading) (*models.SensorReading, error)
}

type airQualityReader struct {
	cfg              *config.Config
	AirQualitySensor AirQualitySensor
}

var ErrorUnknownReading = errors.New("Unknown wanted reading")

func NewAirQualityReader(cfg *config.Config) AirQualityReader {
	airQualitySensor := NewEnviroSensor()
	if cfg.MockHardware {
		airQualitySensor = NewFakeSensor()
	}

	return &airQualityReader{
		cfg:              cfg,
		AirQualitySensor: airQualitySensor,
	}
}

func (a *airQualityReader) ReadSensor(wantedReading Reading) (*models.SensorReading, error) {
	var (
		reading float64
		err     error
	)

	switch wantedReading {
	case Light:
		reading, err = a.AirQualitySensor.ReadLight()
	case HazardousGases:
		reading, err = a.AirQualitySensor.ReadHazardousGases()
	case Humidity:
		reading, err = a.AirQualitySensor.ReadHumidity()
	case Pressure:
		reading, err = a.AirQualitySensor.ReadPressure()
	case Temperature:
		reading, err = a.AirQualitySensor.ReadTemperature()
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
