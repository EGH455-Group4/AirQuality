package sensors_test

import (
	"errors"
	"testing"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/models"
	"github.com/ImTheTom/air-quality/sensors"
	"github.com/ImTheTom/air-quality/sensors/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type SensorTestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	cfg *config.Config

	airQualityReader sensors.AirQualityReader

	airQualitySensor *mocks.MockAirQualitySensor
}

func (s *SensorTestSuite) SetupTest() {
	config.SetupConfig("")

	s.cfg = config.GetConfig()

	s.ctrl = gomock.NewController(s.T())

	s.airQualitySensor = mocks.NewMockAirQualitySensor(s.ctrl)

	s.airQualityReader = sensors.NewAirQualityReaderFromSensor(
		s.cfg,
		s.airQualitySensor,
	)
}

func TestSensorTestSuite(t *testing.T) {
	suite.Run(t, new(SensorTestSuite))
}

func (s *SensorTestSuite) TestSensor_NewAirQualityReader() {
	sensor := sensors.NewAirQualityReader(s.cfg)

	assert.NotNil(s.T(), sensor)
}

func (s *SensorTestSuite) TestSensor_ReadSensorUnspecified() {
	reading, err := s.airQualityReader.ReadSensor(sensors.Unspecified)

	assert.Nil(s.T(), reading)
	assert.Equal(s.T(), sensors.ErrorUnknownReading, err)
}

func (s *SensorTestSuite) TestSensor_ReadSensorLight() {
	s.airQualitySensor.EXPECT().ReadLight().Return(
		55.0, nil,
	)

	reading, err := s.airQualityReader.ReadSensor(sensors.Light)

	assert.NotNil(s.T(), reading)
	assert.Equal(s.T(), &models.SensorReading{
		Reading: 55.0,
	}, reading)
	assert.NoError(s.T(), err)
}

func (s *SensorTestSuite) TestSensor_ReadSensorLight_Error() {
	s.airQualitySensor.EXPECT().ReadLight().Return(
		0.0, errors.New("failed reading"),
	)

	reading, err := s.airQualityReader.ReadSensor(sensors.Light)

	assert.Nil(s.T(), reading)
	assert.Error(s.T(), err)
}

func (s *SensorTestSuite) TestSensor_ReadSensorHazardGases() {
	s.airQualitySensor.EXPECT().ReadHazardousGases().Return(
		26.7, nil,
	)

	reading, err := s.airQualityReader.ReadSensor(sensors.HazardousGases)

	assert.NotNil(s.T(), reading)
	assert.Equal(s.T(), &models.SensorReading{
		Reading: 26.7,
	}, reading)
	assert.NoError(s.T(), err)
}

func (s *SensorTestSuite) TestSensor_ReadSensorHazardGases_Error() {
	s.airQualitySensor.EXPECT().ReadHazardousGases().Return(
		0.0, errors.New("failed reading"),
	)

	reading, err := s.airQualityReader.ReadSensor(sensors.HazardousGases)

	assert.Nil(s.T(), reading)
	assert.Error(s.T(), err)
}

func (s *SensorTestSuite) TestSensor_ReadSensorHumidity() {
	s.airQualitySensor.EXPECT().ReadHumidity().Return(
		77.4, nil,
	)

	reading, err := s.airQualityReader.ReadSensor(sensors.Humidity)

	assert.NotNil(s.T(), reading)
	assert.Equal(s.T(), &models.SensorReading{
		Reading: 77.4,
	}, reading)
	assert.NoError(s.T(), err)
}

func (s *SensorTestSuite) TestSensor_ReadSensorHumidity_Error() {
	s.airQualitySensor.EXPECT().ReadHumidity().Return(
		0.0, errors.New("failed reading"),
	)

	reading, err := s.airQualityReader.ReadSensor(sensors.Humidity)

	assert.Nil(s.T(), reading)
	assert.Error(s.T(), err)
}

func (s *SensorTestSuite) TestSensor_ReadSensorPressure() {
	s.airQualitySensor.EXPECT().ReadPressure().Return(
		5.6, nil,
	)

	reading, err := s.airQualityReader.ReadSensor(sensors.Pressure)

	assert.NotNil(s.T(), reading)
	assert.Equal(s.T(), &models.SensorReading{
		Reading: 5.6,
	}, reading)
	assert.NoError(s.T(), err)
}

func (s *SensorTestSuite) TestSensor_ReadSensorPressure_Error() {
	s.airQualitySensor.EXPECT().ReadPressure().Return(
		0.0, errors.New("failed reading"),
	)

	reading, err := s.airQualityReader.ReadSensor(sensors.Pressure)

	assert.Nil(s.T(), reading)
	assert.Error(s.T(), err)
}

func (s *SensorTestSuite) TestSensor_ReadSensorTemperature() {
	s.airQualitySensor.EXPECT().ReadTemperature().Return(
		23.0, nil,
	)

	reading, err := s.airQualityReader.ReadSensor(sensors.Temperature)

	assert.NotNil(s.T(), reading)
	assert.Equal(s.T(), &models.SensorReading{
		Reading: 23.0,
	}, reading)
	assert.NoError(s.T(), err)
}

func (s *SensorTestSuite) TestSensor_ReadSensorTemperature_Error() {
	s.airQualitySensor.EXPECT().ReadTemperature().Return(
		0.0, errors.New("failed reading"),
	)

	reading, err := s.airQualityReader.ReadSensor(sensors.Temperature)

	assert.Nil(s.T(), reading)
	assert.Error(s.T(), err)
}
