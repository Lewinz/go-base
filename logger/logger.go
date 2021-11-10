package logger

// Logger stack logger
type Logger interface {
	// standard logger
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	// Gen Trace request id
	GenReqID() string
	TraceID() string

	// options
	SetLevel(level Level)
	WithField(field map[string]interface{}) Logger
}
