package sensors_test

import (
	"testing"

	"github.com/ImTheTom/air-quality/sensors"
	"github.com/stretchr/testify/assert"
)

func TestNewFakeSensor(t *testing.T) {
	fk := sensors.NewFakeSensor()

	assert.NotNil(t, fk)
}

func TestReadLight(t *testing.T) {
	fk := sensors.NewFakeSensor()

	readingLight, err := fk.ReadLight()

	if err == nil {
		assert.True(t, readingLight >= 10)
		assert.NoError(t, err)
	} else {
		assert.Equal(t, readingLight, 0.0)
		assert.Error(t, err)
	}
}

func TestReadHazardousGases(t *testing.T) {
	fk := sensors.NewFakeSensor()

	readingHazardousGases, err := fk.ReadHazardousGases()

	if err == nil {
		assert.True(t, readingHazardousGases >= 10)
		assert.NoError(t, err)
	} else {
		assert.Equal(t, readingHazardousGases, 0.0)
		assert.Error(t, err)
	}
}

func TestReadHumidity(t *testing.T) {
	fk := sensors.NewFakeSensor()

	readingHumidity, err := fk.ReadHumidity()

	if err == nil {
		assert.True(t, readingHumidity >= 10)
		assert.NoError(t, err)
	} else {
		assert.Equal(t, readingHumidity, 0.0)
		assert.Error(t, err)
	}
}

func TestReadPressure(t *testing.T) {
	fk := sensors.NewFakeSensor()

	readingPressure, err := fk.ReadPressure()

	if err == nil {
		assert.True(t, readingPressure >= 10)
		assert.NoError(t, err)
	} else {
		assert.Equal(t, readingPressure, 0.0)
		assert.Error(t, err)
	}
}

func TestReadTemperature(t *testing.T) {
	fk := sensors.NewFakeSensor()

	readingTemperature, err := fk.ReadTemperature()

	if err == nil {
		assert.True(t, readingTemperature >= 10)
		assert.NoError(t, err)
	} else {
		assert.Equal(t, readingTemperature, 0.0)
		assert.Error(t, err)
	}
}
