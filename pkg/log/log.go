package log

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Dir = "logs/"
var Level = "trace"

func Init() error {
	_, err := os.Stat(Dir)
	if os.IsNotExist(err) {
		os.MkdirAll(Dir, os.ModeDir|os.ModePerm)
	}
	name := time.Now().Format("2006-01-02T15-04-05")
	path := Dir + name + ".log"
	switch Level {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	}
	//writer := os.Stdout
	writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755) //只写、创建、追加
	if err != nil {
		return err
	}
	logrus.SetOutput(io.MultiWriter(writer))
	return nil
}
