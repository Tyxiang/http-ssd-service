package ssd

import (
	"errors"
	"os"
)

var buffer string
var Dir = "ssds/"

func Init() error {
	_, err := os.Stat(Dir)
	if os.IsNotExist(err) {
		os.MkdirAll(Dir, os.ModeDir|os.ModePerm)
	}
	err = Load("last")
	if err != nil {
		err = errors.New("load last ssd failed")
		return err
	}
	return nil
}
