package models

import "time"

type AirQuality struct {
	Sensors     *Sensors  `json:"sensors"`
	CurrentTime time.Time `json:"current_time"`
	Errors      []string  `json:"errors"`
}

type Sensors struct {
	Light          *SensorReading `json:"light"`
	HazardousGases *SensorReading `json:"hazardous_gases"`
	Humidity       *SensorReading `json:"humidity"`
	Pressure       *SensorReading `json:"pressure"`
	Temperature    *SensorReading `json:"temperature"`
}

type SensorReading struct {
	Reading float64 `json:"reading"`
}

type GeneralResponse struct{}
