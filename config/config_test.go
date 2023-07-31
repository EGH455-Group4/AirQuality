package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigSetup_CreatesDefaultIfMissing(t *testing.T) {
	SetupConfig("")

	assert.Equal(t, &Config{
		MockHardware:      false,
		Port:              ":8050",
		SensorReadSeconds: 5,
		ParallelRead:      false,
		StartOnBoot:       false,
	}, GetConfig())
}

func TestConfigSetup_HandlesJsonFileCorrectly(t *testing.T) {
	SetupConfig("./testing/test.json")

	assert.Equal(t, &Config{
		MockHardware:      true,
		Port:              ":9090",
		SensorReadSeconds: 60,
		ParallelRead:      false,
		StartOnBoot:       true,
	}, GetConfig())
}
