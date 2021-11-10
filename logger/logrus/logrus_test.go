package logrus

import (
	"testing"
	"time"
)

func TestLogrus(t *testing.T) {
	logger := New(GenReqID())

	_ = NewEmptyLoggerWithFields(logger.fields)

	SetLogLevel("debug", logger)
	SetOutput("", logger)

	logger.WithFields(logger.fields)
	logger.Info("test logrus info")
	logger.Debug("test logrus debug")
	logger.Error("test logrus error")
	logger.Warn("test logrus warn")
	logger.Fatal("test logrus fatal")
	logger.Panic("test logrus panic")

	logger.Infof("test logrus infof")
	logger.Debugf("test logrus debugf")
	logger.Errorf("test logrus errorf")
	logger.Warnf("test logrus warnf")
	logger.Fatalf("test logrus fatalf")
	logger.Panicf("test logrus panicf")

	// wait async
	time.Sleep(5 * time.Second)
}
