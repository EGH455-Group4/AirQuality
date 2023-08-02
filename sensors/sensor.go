package sensors

import (
	"errors"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/models"
)

//go:generate mockgen -destination mocks/sensor.mock.go -package mocks -source sensor.go

type AirQualityReader interface {
	ReadSensor(wantedReading Reading) *models.SensorReading
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

func NewAirQualityReaderFromSensor(cfg *config.Config, sensor AirQualitySensor) AirQualityReader {
	return &airQualityReader{
		cfg:              cfg,
		AirQualitySensor: sensor,
	}
}

func (a *airQualityReader) ReadSensor(wantedReading Reading) *models.SensorReading {
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
		err = ErrorUnknownReading
	}

	if err != nil {
		return &models.SensorReading{
			Reading: reading,
			Error:   err.Error(),
		}
	}

	return &models.SensorReading{
		Reading: reading,
	}
}
