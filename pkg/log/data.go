package log

import (
	"github.com/sirupsen/logrus"
)

func Panic(args interface{}) {
	logrus.Panic(args)
}
func Fatal(args interface{}) {
	logrus.Fatal(args)
}
func Error(args interface{}) {
	logrus.Error(args)
}
func Warn(args interface{}) {
	logrus.Warn(args)
}
func Info(args interface{}) {
	logrus.Info(args)
}
func Debug(args interface{}) {
	logrus.Debug(args)
}
func Trace(args interface{}) {
	logrus.Trace(args)
}
