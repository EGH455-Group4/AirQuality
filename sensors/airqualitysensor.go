package sensors

type AirQualitySensor interface {
	ReadLight() (float64, error)
	ReadHazardousGases() (float64, error)
	ReadHumidity() (float64, error)
	ReadPressure() (float64, error)
	ReadTemperature() (float64, error)
}

type Reading int

const (
	Unspecified Reading = iota
	Light
	HazardousGases
	Humidity
	Pressure
	Temperature
)

var ReadingToStringMap = map[Reading]string{
	Unspecified:    "unspecified",
	Light:          "light",
	HazardousGases: "hazardous gases",
	Humidity:       "humidity",
	Pressure:       "pressure",
	Temperature:    "temperature",
}
