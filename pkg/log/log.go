package log

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var dir = "logs/"

func Init() (err error) {
	path := dir + "last" + ".log"
	logrus.SetLevel(logrus.TraceLevel)
	//writer := os.Stdout
	writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755) //只写、创建、追加
	if err != nil {
		return err
	}
	logrus.SetOutput(io.MultiWriter(writer))
	return nil
}
func Save(level string, arg interface{}) (err error) {
	switch level {
	case "trace":
		logrus.Trace(arg)
	case "debug":
		logrus.Debug(arg)
	case "info":
		logrus.Info(arg)
	case "warn":
		logrus.Warn(arg)
	case "error":
		logrus.Error(arg)
	case "fatal":
		logrus.Fatal(arg)
	case "panic":
		logrus.Panic(arg)
	default:
		err := errors.New("no such level")
		return err
	}
	return nil
}
func Read(name string) (data []byte, err error) {
	path := dir + name + ".log"
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func List() (names []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		name := strings.TrimRight(file.Name(), ".log")
		names = append(names, name)
	}
	return names, nil
}
func Remove(name string) (err error) {
	path := dir + name + ".log"
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
