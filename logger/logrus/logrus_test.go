package logrus

import (
	"testing"
	"time"
)

func TestLogrus(t *testing.T) {
	StdLog.Error("test interface")

	// wait async
	time.Sleep(2 * time.Second)
}
