package phuslog

import (
	"testing"
	"time"
)

func TestPhuslog(t *testing.T) {
	StdLog.Error("phuslog logger error")

	time.Sleep(2 * time.Second)
}
