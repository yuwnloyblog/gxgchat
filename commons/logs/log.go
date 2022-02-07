package logs

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/yuwnloyblog/gxgchat/commons/configures"
)

// var LogOut io.Writer
// var ConsoleOut bool

var infoLogger *logrus.Logger
var errorLogger *logrus.Logger

func InitLogs() {
	initErrorLogger()
	initInfoLogger()
}

func initInfoLogger() {
	infoLogger = logrus.New()
	_, err := rotatelogs.New(
		fmt.Sprintf(`%s/%s.%%Y%%m%%d.log`, configures.Config.Log.LogPath, configures.Config.Log.LogName),
		rotatelogs.WithLinkName(fmt.Sprintf(`%s/%s.log`, configures.Config.Log.LogPath, configures.Config.Log.LogName)),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithRotationSize(512*1024*1024),
	)
	if err != nil {
		log.Printf("init log error: %s", err)
		return
	}

	//infoLogger.SetOutput(writer)
	infoLogger.SetOutput(os.Stdout)
	infoLogger.SetReportCaller(true)
	// callerPrettyfier := func(frame *runtime.Frame) (function string, file string) {
	// 	defer func() {
	// 		recover()
	// 	}()
	// 	pc, fullPath, line, _ := runtime.Caller(10)
	// 	funPc := runtime.FuncForPC(pc)
	// 	var funcVal, fileVal string
	// 	funcVal = funPc.Name()
	// 	if strings.Contains(funcVal, "/") {
	// 		funcVal = string([]byte(funcVal)[strings.LastIndex(funcVal, "/")+1:])
	// 	}
	// 	if strings.Contains(fullPath, "/") {
	// 		fileVal = fmt.Sprintf("%s:%d", string([]byte(fullPath)[strings.LastIndex(fullPath, "/")+1:]), line)
	// 	}
	// 	return funcVal, fileVal
	// }
	/*
		if types.SystemEnvDev == setting.Config.Env {
			logrus.SetFormatter(&logrus.TextFormatter{
				ForceColors:      true,
				FullTimestamp:    true,
				CallerPrettyfier: callerPrettyfier,
			})
			logrus.SetOutput(os.Stdout)
		} else {
			logrus.SetFormatter(&logrus.JSONFormatter{
				CallerPrettyfier: callerPrettyfier,
			})
		}
	*/
	// infoLogger.SetFormatter(&logrus.JSONFormatter{
	// 	CallerPrettyfier: callerPrettyfier,
	// })
	infoLogger.SetFormatter(&LogFormatter{})
	infoLogger.SetLevel(logrus.DebugLevel)
}

type LogFormatter struct {
}

func (m *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("20060102150405.123")
	newLog := fmt.Sprintf("[%s] %s\n", timestamp, entry.Message)
	b.WriteString(newLog)
	return b.Bytes(), nil
}

func initErrorLogger() {
	errorLogger = logrus.New()
	//writer
	_, err := rotatelogs.New(
		fmt.Sprintf(`%s/%s.%%Y%%m%%d.log`, configures.Config.Log.LogPath, configures.Config.Log.LogName+"_err"),
		rotatelogs.WithLinkName(fmt.Sprintf(`%s/%s.log`, configures.Config.Log.LogPath, configures.Config.Log.LogName+"_err")),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithRotationSize(512*1024*1024),
	)
	if err != nil {
		log.Printf("init log error: %s", err)
		return
	}

	//errorLogger.SetOutput(writer)
	errorLogger.SetOutput(os.Stdout)
	errorLogger.SetReportCaller(true)
	errorLogger.SetFormatter(&LogFormatter{})
	errorLogger.SetLevel(logrus.WarnLevel)
}

func Panic(f interface{}, v ...interface{}) {
	errorLogger.Panic(f, v)
}

func Fata(f interface{}, v ...interface{}) {
	errorLogger.Fatal(f, v)
}

func Error(f interface{}, v ...interface{}) {
	errorLogger.Error(f, v)
}
func Warn(f interface{}, v ...interface{}) {
	errorLogger.Warn(f, v)
}

func Info(f interface{}, v ...interface{}) {
	infoLogger.Info(f, v)
}

func Debug(f interface{}, v ...interface{}) {
	infoLogger.Debug(f, v)
}

func Trace(f interface{}, v ...interface{}) {
	infoLogger.Trace(f, v)
}
