package config

import "os"

var buffer string
var Dir = "configs/"
var Warn error

func Init() error {
	_, err := os.Stat(Dir)
	if os.IsNotExist(err) {
		os.MkdirAll(Dir, os.ModeDir|os.ModePerm)
	}
	err = Load("last")
	if err != nil {
		Warn = err
		err = Load("default")
		if err != nil {
			return err
		}
	}
	return err
}
