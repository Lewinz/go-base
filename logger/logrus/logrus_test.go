package logrus

import (
	"testing"
	"time"

	"github.com/Lewinz/golang_utils/logger"
)

func TestLogrus(t *testing.T) {
	str := &LoggerStruct{
		logger: StdLog,
	}
	str.logger.Error("test interface")

	// wait async
	time.Sleep(2 * time.Second)
}

type LoggerStruct struct {
	logger logger.Logger
}
