package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerSetup(t *testing.T) {
	result := LoggerSetup()

	assert.True(t, result)
}
