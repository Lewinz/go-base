package zerolog

import (
	"sync"

	log "github.com/rs/zerolog"
)

// Logger zerolog implementation
type Logger struct {
	mu     sync.RWMutex
	Logger log.Logger
}

// Debug ..
func (logger *Logger) Debug(args ...interface{}) {
	// logger.Logger.Debug().Msg(args...)
}

// Info ..
func (logger *Logger) Info(args ...interface{}) {
	// logger.Logger.Info().Msgs(args...)
}

// Warn ..
func (logger *Logger) Warn(args ...interface{}) {
	// logger.Logger.Warn().Msgs(args...)
}

// Error ..
func (logger *Logger) Error(args ...interface{}) {
	// logger.Logger.Error().Msgs(args...)
}

// Fatal ..
func (logger *Logger) Fatal(args ...interface{}) {
	// logger.Logger.Fatal().Msgs(args...)
}

// Panic ..
func (logger *Logger) Panic(args ...interface{}) {
	// logger.Logger.Panic().Msgs(args...)
}

// Debugf ..
func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.Logger.Debug().Msgf(format, args...)
}

// Infof ..
func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.Logger.Info().Msgf(format, args...)
}

// Warnf ..
func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.Logger.Warn().Msgf(format, args...)
}

// Errorf ..
func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.Logger.Error().Msgf(format, args...)
}

// Fatalf ..
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.Logger.Fatal().Msgf(format, args...)
}

// Panicf ..
func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.Logger.Panic().Msgf(format, args...)
}

// ReqID ..
func (logger *Logger) ReqID() string {
	return ""
}

// SetLogLevel ..
func SetLogLevel(level string, logger *Logger) {
	switch level {
	case "debug":
		logger.SetLevel(log.DebugLevel)
	case "info":
		logger.SetLevel(log.InfoLevel)
	case "warn":
		logger.SetLevel(log.WarnLevel)
	case "error":
		logger.SetLevel(log.ErrorLevel)
	case "fatal":
		logger.SetLevel(log.FatalLevel)
	case "panic":
		logger.SetLevel(log.PanicLevel)
	default:
		logger.SetLevel(log.InfoLevel)
	}
}

// SetLevel ..
func (logger *Logger) SetLevel(level log.Level) {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	logger.Logger.Level(level)
}
