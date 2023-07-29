package sensors

type EnviroSensor struct{}

func NewEnviroSensor() AirQualitySensor {
	return &EnviroSensor{}
}

func (*EnviroSensor) ReadLight() (float64, error) {
	return 0, nil
}

func (*EnviroSensor) ReadHazardousGases() (float64, error) {
	return 0, nil
}

func (*EnviroSensor) ReadHumidity() (float64, error) {
	return 0, nil
}

func (*EnviroSensor) ReadPressure() (float64, error) {
	return 0, nil
}

func (*EnviroSensor) ReadTemperature() (float64, error) {
	return 0, nil
}
