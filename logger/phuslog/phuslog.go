package phuslog

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/phuslu/log"
)

var (
	// RequestIDKey RequestID fields
	RequestIDKey = "x-reqid"
)

// logger constant
var (
	StdLog *Logger
)

func init() {
	StdLog = NewEmptyLogger()
	SetAsyncFileOutput("phusStd.log", 1000, StdLog)
}

// Logger phuslog implementation
type Logger struct {
	mu     sync.RWMutex
	Logger *log.Logger
	reqID  string
}

// New ..
func New(args ...interface{}) *Logger {
	reqID := ""
	if len(args) > 0 && args[0] != nil {
		a := args[0]
		switch a.(type) {
		case *log.Logger:
			return NewLogger(a.(*log.Logger))
		case string:
			reqID = a.(string)
		}
	}

	logger := NewEmptyLogger()

	if reqID == "" {
		reqID = GetReqID()
	}

	ctx := log.NewContext(nil).Str(RequestIDKey, reqID).Value()
	logger.Logger.Context = ctx
	return logger
}

// NewLogger new logger
func NewLogger(logg *log.Logger) *Logger {
	logger := &Logger{
		Logger: logg,
	}
	return logger
}

// NewEmptyLogger new empty logger
func NewEmptyLogger() *Logger {
	logger := &Logger{
		Logger: &log.Logger{
			Caller:     2,
			TimeField:  "data",
			TimeFormat: "2006-01-02 15:04:05",
		},
	}

	SetOutput("", logger)
	return logger
}

// Debug ..
func (logger *Logger) Debug(args ...interface{}) {
	logger.Logger.Debug().Msgs(args...)
}

// Info ..
func (logger *Logger) Info(args ...interface{}) {
	logger.Logger.Info().Msgs(args...)
}

// Warn ..
func (logger *Logger) Warn(args ...interface{}) {
	logger.Logger.Warn().Msgs(args...)
}

// Error ..
func (logger *Logger) Error(args ...interface{}) {
	logger.Logger.Error().Msgs(args...)
}

// Fatal ..
func (logger *Logger) Fatal(args ...interface{}) {
	logger.Logger.Fatal().Msgs(args...)
}

// Panic ..
func (logger *Logger) Panic(args ...interface{}) {
	logger.Logger.Panic().Msgs(args...)
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
	return logger.reqID
}

// SetLogLevel other package set level
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

// SetLevel set phuslog level
func (logger *Logger) SetLevel(level log.Level) {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	logger.Logger.SetLevel(level)
}

// SetOutput ..
func SetOutput(output string, logger *Logger) {
	writer := &log.IOWriter{
		Writer: os.Stderr,
	}

	switch output {
	case "stdout":
		writer.Writer = os.Stdout
	case "null":
		writer.Writer = ioutil.Discard
	}
	logger.SetOutput(writer)
}

// SetAsyncFileOutput set logger async write
func SetAsyncFileOutput(fileName string, channelSize uint, logger *Logger) {
	writer := &log.AsyncWriter{
		ChannelSize: channelSize,
		Writer: &log.FileWriter{
			Filename: fileName,
			FileMode: 0600,
			// 日志写入文件前确保日志目录被创建
			EnsureFolder: true,
			// 日志文件名时间格式
			TimeFormat: "2006-01-02",
			// 采用服务器本地时间，false 为使用 UTC 时间
			LocalTime: true,
		},
	}

	logger.SetOutput(writer)
}

// SetOutput sets the standard logger output.
func (logger *Logger) SetOutput(out log.Writer) {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	logger.Logger.Writer = out
}

// GetReqID get req id
func GetReqID() string {
	return log.NewXID().String()
}
