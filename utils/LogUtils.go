package utils

import (
	"AliyunStat/settings"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
)

//初始化日志组件
// config logrus log to local filesystem, with file rotation
func init() {
	baseLogPath := path.Join(settings.LogPath, settings.LogFileName)

	writer, err := rotatelogs.New(baseLogPath+".%Y%m%d%H%M", rotatelogs.WithLinkName(baseLogPath), rotatelogs.WithMaxAge(settings.MaxAge), rotatelogs.WithRotationTime(settings.RotationTime))

	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	// 以JSON格式为输出，代替默认的ASCII格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 以Stdout为输出，代替默认的stderr
	//logrus.SetOutput(os.Stdout)
	// 设置日志等级
	//logrus.SetLevel(logrus.WarnLevel)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		//为不同级别设置不同的输出目的
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	})

	logrus.AddHook(lfHook)
}

func LogDebug(log string) {
	logrus.Debug(log)
}

func LogInfo(log string) {
	logrus.Info(log)
}

func LogWarn(log string) {
	logrus.Warn(log)
}

func LogError(log string) {
	logrus.Error(log)
}

func LogFatal(log string) {
	logrus.Fatal(log)
}

func LogPanic(log string) {
	logrus.Panic(log)
}
