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
		errStirng := "load last ssd error, load defualt"
		err = Load("default")
		if err != nil {
			errStirng = errStirng + ". load default ssd error"
			err = errors.New(errStirng)
			buffer = ""
			return err
		}
	}
	return nil
}
