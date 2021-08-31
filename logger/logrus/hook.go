package logrus

// import (
// 	"time"

// 	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
// 	log "github.com/sirupsen/logrus"
// )

// func newLfsHook(logLevel *string, maxRemainCnt uint) log.Hook {
// 	writer, err := rotatelogs.New(
// 		logName+".%Y%m%d%H",
// 		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
// 		rotatelogs.WithLinkName(logName),

// 		// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
// 		rotatelogs.WithRotationTime(time.Hour),

// 		// WithMaxAge和WithRotationCount二者只能设置一个,
// 		// WithMaxAge设置文件清理前的最长保存时间,
// 		// WithRotationCount设置文件清理前最多保存的个数.
// 		//rotatelogs.WithMaxAge(time.Hour*24),
// 		rotatelogs.WithRotationCount(maxRemainCnt),
// 	)

// 	if err != nil {
// 		log.Errorf("config local file system for logger error: %v", err)
// 	}

// 	level, ok := logLevels[*logLevel]

// 	if ok {
// 		log.SetLevel(level)
// 	} else {
// 		log.SetLevel(log.WarnLevel)
// 	}

// 	lfsHook := lfshook.NewHook(lfshook.WriterMap{
// 		log.DebugLevel: writer,
// 		log.InfoLevel:  writer,
// 		log.WarnLevel:  writer,
// 		log.ErrorLevel: writer,
// 		log.FatalLevel: writer,
// 		log.PanicLevel: writer,
// 	}, &log.TextFormatter{DisableColors: true})

// 	return lfsHook
// }
