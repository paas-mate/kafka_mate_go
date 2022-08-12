package service

import (
	"github.com/stretchr/testify/assert"
	"kafka_mate_go/pkg/test"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	test.SetupKafka()
	err := HealthCheck()
	assert.Nil(t, err)
}
