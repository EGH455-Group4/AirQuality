package service_test

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/models"
	"github.com/ImTheTom/air-quality/sensors"
	"github.com/ImTheTom/air-quality/sensors/mocks"
	"github.com/ImTheTom/air-quality/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

var expectedModelForReadSensors = &models.Sensors{
	Light: &models.SensorReading{
		Reading: 55.0,
	},
	HazardousGases: &models.SensorReading{
		Reading: 23.6,
	},
	Humidity: nil,
	Pressure: &models.SensorReading{
		Reading: 77.2,
	},
	Temperature: &models.SensorReading{
		Reading: 95.2,
	},
}

type AirQualityTestSuite struct {
	suite.Suite

	ctrl *gomock.Controller

	cfg *config.Config

	airQualityReader *mocks.MockAirQualityReader

	airQualityService service.AirQualityService
}

func (s *AirQualityTestSuite) SetupTest() {
	config.SetupConfig("")

	s.cfg = config.GetConfig()

	s.ctrl = gomock.NewController(s.T())

	s.airQualityReader = mocks.NewMockAirQualityReader(s.ctrl)

	s.airQualityService = service.NewAirQualityService(
		s.cfg,
		s.airQualityReader,
	)
}

func TestAirQualityTestSuite(t *testing.T) {
	suite.Run(t, new(AirQualityTestSuite))
}

func (s *AirQualityTestSuite) TestSensor_GetAirQuality() {
	airQuality := s.airQualityService.GetAirQuality()

	assert.NotNil(s.T(), airQuality)
	assert.Empty(s.T(), airQuality.Errors)
	assert.NotEqual(s.T(), time.Time{}, airQuality.CurrentTime)
}

func (s *AirQualityTestSuite) TestSensor_Start() {
	s.airQualityService.Start()
}

func (s *AirQualityTestSuite) TestSensor_Stop() {
	s.airQualityService.Stop()
}

func (s *AirQualityTestSuite) TestSensor_SingleRead() {
	s.MockExpectedCallsForRead()

	airQuality := s.airQualityService.SingleRead()

	assert.NotNil(s.T(), airQuality)
	assert.Equal(s.T(), 1, len(airQuality.Errors))
	assert.Equal(s.T(), expectedModelForReadSensors, airQuality.Sensors)
}

func (s *AirQualityTestSuite) TestSensor_SingleReadParallel() {
	s.MockExpectedCallsForRead()

	s.cfg.ParallelRead = true

	airQuality := s.airQualityService.SingleRead()

	assert.NotNil(s.T(), airQuality)
	assert.Equal(s.T(), 1, len(airQuality.Errors))
	assert.Equal(s.T(), expectedModelForReadSensors, airQuality.Sensors)
}

func (s *AirQualityTestSuite) TestSensor_RunReadSensors() {
	s.MockExpectedCallsForRead()

	testWg := sync.WaitGroup{}

	readSensorQuitChan := make(chan bool)

	s.airQualityService.Start()

	go s.airQualityService.RunReadSensors(readSensorQuitChan, &testWg)

	time.Sleep(1 * time.Second)

	readSensorQuitChan <- true

	testWg.Wait()

	airQuality := s.airQualityService.GetAirQuality()

	assert.NotNil(s.T(), airQuality)
	assert.Equal(s.T(), 1, len(airQuality.Errors))
	assert.Equal(s.T(), expectedModelForReadSensors, airQuality.Sensors)
}

func (s *AirQualityTestSuite) MockExpectedCallsForRead() {
	s.airQualityReader.EXPECT().ReadSensor(sensors.Light).Return(
		&models.SensorReading{
			Reading: 55.0,
		}, nil,
	)

	s.airQualityReader.EXPECT().ReadSensor(sensors.HazardousGases).Return(
		&models.SensorReading{
			Reading: 23.6,
		}, nil,
	)

	s.airQualityReader.EXPECT().ReadSensor(sensors.Humidity).Return(
		nil, errors.New("test error"),
	)

	s.airQualityReader.EXPECT().ReadSensor(sensors.Pressure).Return(
		&models.SensorReading{
			Reading: 77.2,
		}, nil,
	)

	s.airQualityReader.EXPECT().ReadSensor(sensors.Temperature).Return(
		&models.SensorReading{
			Reading: 95.2,
		}, nil,
	)
}
