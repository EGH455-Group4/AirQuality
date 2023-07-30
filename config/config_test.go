package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigSetup_CreatesDefaultIfMissing(t *testing.T) {
	SetupConfig("")

	assert.Equal(t, &Config{
		MockHardware:      false,
		Address:           ":8050",
		SensorReadSeconds: 5,
		ParallelRead:      false,
	}, GetConfig())
}

func TestConfigSetup_HandlesJsonFileCorrectly(t *testing.T) {
	SetupConfig("./testing/test.json")

	assert.Equal(t, &Config{
		MockHardware:      true,
		Address:           ":9090",
		SensorReadSeconds: 60,
		ParallelRead:      false,
	}, GetConfig())
}
