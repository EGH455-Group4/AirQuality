package sensors

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrorReadingValue = errors.New("failed reading value")

type FakeSensor struct{}

func NewFakeSensor() AirQualitySensor {
	return &FakeSensor{}
}

func (*FakeSensor) ReadLight() (float64, error) {
	return readRandomValue(Light)
}

func (*FakeSensor) ReadHazardousGases() (float64, error) {
	return readRandomValue(HazardousGases)
}

func (*FakeSensor) ReadHumidity() (float64, error) {
	return readRandomValue(Humidity)
}

func (*FakeSensor) ReadPressure() (float64, error) {
	return readRandomValue(Pressure)
}

func (*FakeSensor) ReadTemperature() (float64, error) {
	return readRandomValue(Temperature)
}

func readRandomValue(wantedReading Reading) (float64, error) {
	randomValue := rand.Float64() * 100

	if randomValue < 10 {
		return 0, fmt.Errorf("%s failed reading it's value", ReadingToStringMap[wantedReading])
	}

	return randomValue, nil
}
