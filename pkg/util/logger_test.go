package util

import (
	"testing"
)

func TestLogger(t *testing.T) {
	logger := Logger()
	if logger == nil {
		t.Errorf("logger is null")
	}
}
