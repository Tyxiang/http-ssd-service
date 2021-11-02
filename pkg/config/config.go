package config

import (
	"errors"
	"os"
)

var buffer string
var Dir = "configs/"

func Init() error {
	_, err := os.Stat(Dir)
	if os.IsNotExist(err) {
		os.MkdirAll(Dir, os.ModeDir|os.ModePerm)
	}
	err = Load("last")
	if err != nil {
		err = errors.New("load last config failed")
		return err
	}
	return nil
}
