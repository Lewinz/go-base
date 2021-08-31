package logrus

import (
	"encoding/base64"
	"encoding/binary"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	// RequestIDKey RequestID fields
	RequestIDKey = "x-reqid"
	pid          = uint32(time.Now().UnixNano() % 4294967291)

	logQueue = make(chan *LogMessage, 1000)
)

// logger constant
var (
	StdLog *Logger
)

func init() {
	// async logger
	asyncLoggerWrite()
	// Std logger
	StdLog = NewLogger()
}

// Logger logrus realize
type Logger struct {
	mu       sync.RWMutex
	out      io.Writer
	LogEntry *log.Entry
	fields   map[string]interface{}
	reqID    string
}

// WithFields ..
func (logger *Logger) WithFields(fields map[string]interface{}) *log.Entry {
	f := make(log.Fields)
	for k, v := range fields {
		f[k] = v
	}

	logger.LogEntry = logger.LogEntry.WithFields(f)
	return logger.LogEntry
}

// SetLevel ..
func (logger *Logger) SetLevel(level log.Level) {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	logger.LogEntry.Logger.SetLevel(level)
}

// SetOutput sets the standard logger output.
func (logger *Logger) SetOutput(out io.Writer) {
	logger.mu.Lock()
	defer logger.mu.Unlock()

	logger.LogEntry.Logger.Out = out
}

// GenReqID ..
func GenReqID() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}

// New this function is recommended to create a new logger object
func New(o ...interface{}) *Logger {
	var reqID = ""
	if len(o) > 0 && o[0] != nil {
		a := o[0]

		switch a.(type) {
		case *Logger:
			return a.(*Logger)
		case *log.Entry:
			return NewLogger(a.(*log.Entry))
		case string:
			reqID = a.(string)
		}
	}

	if len(reqID) == 0 {
		reqID = GenReqID()
	}

	l := NewEmptyLoggerWithFields(map[string]interface{}{RequestIDKey: reqID})
	l.reqID = reqID
	return l
}

// NewLogger ..
func NewLogger(entrys ...*log.Entry) *Logger {
	logger := log.New()
	l := &Logger{
		LogEntry: log.NewEntry(logger),
		fields:   make(map[string]interface{}),
	}

	if len(entrys) > 0 {
		l.LogEntry = entrys[0]
	}
	SetOutputFile("log.log", l)
	SetFormat("json", l)
	return l
}

// LoggerFields ..
type LoggerFields map[string]interface{}

// NewEmptyLoggerWithFields ..
func NewEmptyLoggerWithFields(fields map[string]interface{}) *Logger {
	l := NewLogger()
	l.fields = fields
	return l
}

// DecorateLog ..
func DecorateLog(logger *log.Entry) *log.Entry {
	var (
		fileName, funcName string
	)
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		fileName = "???"
		funcName = "???"
		line = 0
	} else {
		funcName = runtime.FuncForPC(pc).Name()
		fileSlice := strings.Split(file, path.Dir(funcName))
		fileName = filepath.Join(path.Dir(funcName), fileSlice[len(fileSlice)-1]) + ":" + strconv.Itoa(line)
	}
	return logger.WithField("file", fileName).WithField("func", funcName)
}

// Hook ..
type Hook interface {
	Levels() []log.Level
	Fire(*log.Entry) error
}

// AddHook ..
func (logger *Logger) AddHook(hook Hook) *Logger {
	logger.LogEntry.Logger.AddHook(hook)
	return logger
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

// SetOutput ..
func SetOutput(output string, logger *Logger) {
	switch output {
	case "stderr":
		logger.SetOutput(os.Stderr)
	case "stdout":
		logger.SetOutput(os.Stdout)
	case "null":
		logger.SetOutput(ioutil.Discard)
	default:
		logger.SetOutput(os.Stderr)
	}
}

// SetOutputFile set logger file
func SetOutputFile(fileName string, logger *Logger) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		logger.Errorf("set output file err(%v),file(%v)", err, fileName)
	}
	logger.SetOutput(io.MultiWriter(file))
}

// SetFormat ..
func SetFormat(format string, logger *Logger) {
	switch format {
	case "json":
		logger.LogEntry.Logger.Formatter = &log.JSONFormatter{
			DisableTimestamp: true,
		}
	case "text":
		logger.LogEntry.Logger.Formatter = &log.TextFormatter{
			DisableTimestamp: true,
		}
	default:
		logger.LogEntry.Logger.Formatter = &log.JSONFormatter{
			DisableTimestamp: true,
		}
	}
}

// Debug ..
func (logger *Logger) Debug(args ...interface{}) {
	appendLog(&LogMessage{
		Level: "debug",
		Entry: initLogger(logger),
		Data:  args,
	})
}

// Info ..
func (logger *Logger) Info(args ...interface{}) {
	appendLog(&LogMessage{
		Level: "info",
		Entry: initLogger(logger),
		Data:  args,
	})
}

// Warn ..
func (logger *Logger) Warn(args ...interface{}) {
	appendLog(&LogMessage{
		Level: "warn",
		Entry: initLogger(logger),
		Data:  args,
	})
}

// Error ..
func (logger *Logger) Error(args ...interface{}) {
	appendLog(&LogMessage{
		Level: "error",
		Entry: initLogger(logger),
		Data:  args,
	})
}

// Fatal ..
func (logger *Logger) Fatal(args ...interface{}) {
	appendLog(&LogMessage{
		Level: "fatal",
		Entry: initLogger(logger),
		Data:  args,
	})
}

// Panic ..
func (logger *Logger) Panic(args ...interface{}) {
	appendLog(&LogMessage{
		Level: "panic",
		Entry: initLogger(logger),
		Data:  args,
	})
}

// Debugf ..
func (logger *Logger) Debugf(format string, args ...interface{}) {
	appendLog(&LogMessage{
		Level:  "debugf",
		Format: format,
		Entry:  initLogger(logger),
		Data:   args,
	})
}

// Infof ..
func (logger *Logger) Infof(format string, args ...interface{}) {
	appendLog(&LogMessage{
		Level:  "infof",
		Format: format,
		Entry:  initLogger(logger),
		Data:   args,
	})
}

// Warnf ..
func (logger *Logger) Warnf(format string, args ...interface{}) {
	appendLog(&LogMessage{
		Level:  "warnf",
		Format: format,
		Entry:  initLogger(logger),
		Data:   args,
	})
}

// Errorf ..
func (logger *Logger) Errorf(format string, args ...interface{}) {
	appendLog(&LogMessage{
		Level:  "errorf",
		Format: format,
		Entry:  initLogger(logger),
		Data:   args,
	})
}

// Fatalf ..
func (logger *Logger) Fatalf(format string, args ...interface{}) {
	appendLog(&LogMessage{
		Level:  "fatalf",
		Format: format,
		Entry:  initLogger(logger),
		Data:   args,
	})
}

// Panicf ..
func (logger *Logger) Panicf(format string, args ...interface{}) {
	appendLog(&LogMessage{
		Level:  "panicf",
		Format: format,
		Entry:  initLogger(logger),
		Data:   args,
	})
}

// ReqID ..
func (logger *Logger) ReqID() string {
	return logger.reqID
}

// initLogger add common field
func initLogger(logger *Logger) *log.Entry {
	entry := logger.LogEntry
	f := logger.fields
	return DecorateLog(entry.WithFields(log.Fields(f)).WithFields(log.Fields{"timedate": time.Now()}))
}

// Xput ..
func (logger *Logger) Xput(logs []string) {
	if xLog, exists := logger.LogEntry.Data["X-Log"]; exists {
		if ll, ok := xLog.([]string); ok {
			ll = append(ll, logs...)
		}
	} else {
		logger.LogEntry.Data["X-Log"] = logs
	}
}

// LogMessage async logger data
type LogMessage struct {
	Level  string
	Format string
	Data   []interface{}
	Entry  *log.Entry
}

func appendLog(message *LogMessage) {
	if len(logQueue) < cap(logQueue) {
		logQueue <- message
	}
}

// asyncLoggerWrite async logger
func asyncLoggerWrite() {
	go func() {
		goroutineCount := make(chan struct{}, 10)

		for message := range logQueue {
			if float64(len(logQueue))/float64(cap(logQueue)) > 0.7 {
				goroutineCount <- struct{}{}
				go func() {
					defer func() {
						<-goroutineCount
					}()
					levelLoggerPrint(message)
				}()
			} else {
				levelLoggerPrint(message)
			}
		}
	}()
}

// levelLoggerPrint print
func levelLoggerPrint(message *LogMessage) {
	switch message.Level {
	case "debug":
		message.Entry.Debug(message.Data...)
	case "info":
		message.Entry.Info(message.Data...)
	case "warn":
		message.Entry.Warn(message.Data...)
	case "error":
		message.Entry.Error(message.Data...)
	case "fatal":
		message.Entry.Fatal(message.Data...)
	case "panic":
		message.Entry.Panic(message.Data...)
	case "debugf":
		message.Entry.Debugf(message.Format, message.Data...)
	case "infof":
		message.Entry.Infof(message.Format, message.Data...)
	case "warnf":
		message.Entry.Warnf(message.Format, message.Data...)
	case "errorf":
		message.Entry.Errorf(message.Format, message.Data...)
	case "fatalf":
		message.Entry.Fatalf(message.Format, message.Data...)
	case "panicf":
		message.Entry.Panicf(message.Format, message.Data...)
	}
}
