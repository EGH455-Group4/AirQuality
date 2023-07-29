package service

import (
	"time"

	"github.com/ImTheTom/air-quality/config"
	"github.com/ImTheTom/air-quality/models"
)

type AirQualityService struct {
	cfg     *config.Config
	Sensors *models.Sensors
	Errors  []string
	Running bool
}

func NewAirQualityService(cfg *config.Config) *AirQualityService {
	return &AirQualityService{
		cfg:     cfg,
		Sensors: &models.Sensors{},
		Errors:  []string{},
		Running: false,
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
}
