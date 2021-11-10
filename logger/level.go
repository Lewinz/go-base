package logger

import "github.com/phuslu/log"

// Level logger stack level
type Level string

// Level enum
const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelFatal Level = "fatal"
	LevelPanic Level = "panic"
)

// ToPhuslogLevel covert phuslog level
func (e Level) ToPhuslogLevel() log.Level {
	switch e {
	case LevelDebug:
		return log.DebugLevel
	case LevelInfo:
		return log.InfoLevel
	case LevelWarn:
		return log.WarnLevel
	case LevelError:
		return log.ErrorLevel
	case LevelFatal:
		return log.FatalLevel
	case LevelPanic:
		return log.PanicLevel
	default:
		return log.InfoLevel
	}
}
