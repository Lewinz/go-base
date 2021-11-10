package phuslog

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/Lewinz/golang_utils/logger"
	"github.com/phuslu/log"
)

var (
	// RequestIDKey RequestID fields
	RequestIDKey = "x-req-id"

	// DefaultFieldTimeFormat logger field time format
	DefaultFieldTimeFormat = "2006-01-02 15:04:05.000000"

	// DefaultFileTimeFormat logger file name format
	DefaultFileTimeFormat = "2006-01-02"
)

// ----------------------------------------------------------------

// logger constant
var (
	StdLog *Log
)

func init() {
	StdLog = NewLogger()
	SetAsyncFileOutput(StdLog, "raccoon_std.log", 1000)
}

// ----------------------------------------------------------------

// Log phuslog implementation
type Log struct {
	mu     sync.RWMutex
	Logger *log.Logger
	reqID  string
}

// New ..
func New(args ...interface{}) *Log {
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

	l := NewLogger()

	if reqID == "" {
		reqID = l.GenReqID()
	}

	l.Logger.Context = log.NewContext(nil).Str(RequestIDKey, reqID).Value()

	return l
}

// NewLogger new logger
func NewLogger(l ...*log.Logger) *Log {
	res := &Log{
		Logger: &log.Logger{
			Caller:     2,
			TimeField:  "data",
			TimeFormat: DefaultFieldTimeFormat,
		},
	}

	if len(l) > 0 {
		res.Logger = l[0]
	}

	setOutput(res, "")
	return res
}

// ---------------------------------------------------------------- interface func

// Debug pritn log msg
func (l *Log) Debug(args ...interface{}) {
	l.Logger.Debug().Msgs(args...)
}

// Info pritn log msg
func (l *Log) Info(args ...interface{}) {
	l.Logger.Info().Msgs(args...)
}

// Warn pritn log msg
func (l *Log) Warn(args ...interface{}) {
	l.Logger.Warn().Msgs(args...)
}

// Error pritn log msg
func (l *Log) Error(args ...interface{}) {
	l.Logger.Error().Msgs(args...)
}

// Fatal pritn log msg
func (l *Log) Fatal(args ...interface{}) {
	l.Logger.Fatal().Msgs(args...)
}

// Panic pritn log msg
func (l *Log) Panic(args ...interface{}) {
	l.Logger.Panic().Msgs(args...)
}

// Debugf pritn log msg
func (l *Log) Debugf(format string, args ...interface{}) {
	l.Logger.Debug().Msgf(format, args...)
}

// Infof pritn log msg
func (l *Log) Infof(format string, args ...interface{}) {
	l.Logger.Info().Msgf(format, args...)
}

// Warnf pritn log msg
func (l *Log) Warnf(format string, args ...interface{}) {
	l.Logger.Warn().Msgf(format, args...)
}

// Errorf pritn log msg
func (l *Log) Errorf(format string, args ...interface{}) {
	l.Logger.Error().Msgf(format, args...)
}

// Fatalf pritn log msg
func (l *Log) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatal().Msgf(format, args...)
}

// Panicf pritn log msg
func (l *Log) Panicf(format string, args ...interface{}) {
	l.Logger.Panic().Msgf(format, args...)
}

// GenReqID generate request id
func (l *Log) GenReqID() string {
	return log.NewXID().String()
}

// TraceID trace request id
func (l *Log) TraceID() string {
	return l.reqID
}

// SetLevel set phuslog level
func (l *Log) SetLevel(level logger.Level) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Logger.SetLevel(level.ToPhuslogLevel())
}

// WithField set field
func (l *Log) WithField(field map[string]interface{}) logger.Logger {
	entry := l.Logger.Debug()
	for k, v := range field {
		switch v.(type) {
		case string:
			entry = entry.Str(k, v.(string))
		case int:
			entry = entry.Int(k, v.(int))
		}
	}
	l.Logger.Context = entry.Value()
	return l
}

// ---------------------------------------------------------------- implementation add func

// SetOutput sets the standard logger output.
func (l *Log) setOutput(out log.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Logger.Writer = out
}

// ---------------------------------------------------------------- package func

// SetOutput ..
func setOutput(l *Log, output string) {
	w := &log.IOWriter{
		Writer: os.Stderr,
	}

	switch output {
	case "stdout":
		// 缓冲
		w.Writer = os.Stdout
	case "null":
		// 无条件成功输出
		w.Writer = ioutil.Discard
	}

	l.setOutput(w)
}

// SetAsyncFileOutput set logger async write
func SetAsyncFileOutput(l *Log, fileName string, channelSize uint) {
	l.setOutput(&log.AsyncWriter{
		ChannelSize: channelSize,
		Writer: &log.FileWriter{
			Filename: fileName,
			FileMode: 0600,
			// 日志写入文件前确保日志目录被创建
			EnsureFolder: true,
			// 日志文件名时间格式
			TimeFormat: DefaultFileTimeFormat,
			// 采用服务器本地时间，false 为使用 UTC 时间
			LocalTime: true,
		},
	})
}
