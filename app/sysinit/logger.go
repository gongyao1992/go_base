package sysinit

import (
	log "github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"io"
	"micro/upperspective/app/config"
	"os"
)

var (
	lInfo *log.Logger
	dbLogger *log.Logger
)

// MyHook 发送钉钉信息的hook
type MyHook struct {
}

// Levels 只定义 error 和 panic 等级的日志,其他日志等级不会触发 hook
func (h *MyHook) Levels() []log.Level {
	return []log.Level{
		log.ErrorLevel,
		log.PanicLevel,
	}
}
// Fire 将异常日志写入到指定日志文件中
func (h *MyHook) Fire(entry *log.Entry) error {
	//DingtalkSend(entry.Message)
	return nil
}

func init()  {
	wd, _ := os.Getwd()

	// 设置普通log
	lInfo = log.New()
	lInfo.SetFormatter(&log.JSONFormatter{})
	loggerDir := wd + config.Config.Log.Dir + "vip.log"
	loggerWriter := getWriter(loggerDir)
	lInfo.SetOutput(loggerWriter)
	// 添加钩子
	lInfo.AddHook(&MyHook{})

	// 设置db log
	dbLoggerDir := wd + config.Config.Log.DbDir + "db.log"
	dbLogger = log.New()
	dbLogger.SetFormatter(&log.JSONFormatter{})
	//DbLogger.SetLevel(log.InfoLevel)
	//DbLogger.SetReportCaller(true)
	dbLoggerWriter := getWriter(dbLoggerDir)
	dbLogger.SetOutput(dbLoggerWriter)
}

func getWriter(dbLoggerDir string) io.Writer {
	logger:=&lumberjack.Logger{
		LocalTime:  true,
		Filename:   dbLoggerDir,
		MaxSize:    20, // megabytes
		MaxBackups: 5,
		MaxAge:     30,    //days
		Compress:   false, // disabled by default
	}
	writers := []io.Writer{
		logger,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	return fileAndStdoutWriter
}

func VipLogger() *log.Logger {
	return lInfo
}
func DbLogger() *log.Logger {
	return dbLogger
}