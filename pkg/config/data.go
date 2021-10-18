package config

import (
	"errors"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func Add(path string, data []byte) error {
	if path == "" {
		if gjson.GetBytes(bufferBytes, "@this").Exists() {
			err := errors.New("already exist")
			return err
		}
		bufferBytes = data
	}
	if path != "" {
		if gjson.GetBytes(bufferBytes, path).Exists() {
			err := errors.New("already exist")
			return err
		}
		var err error
		bufferBytes, err = sjson.SetRawBytes(bufferBytes, path, data)
		if err != nil {
			return err
		}
	}
	return nil
}
func Get(path string) (interface{}, error) {
	if path == "" {
		path = "@this"
	}
	data := gjson.GetBytes(bufferBytes, path).Value()
	if data == nil {
		err := errors.New("not exist")
		return nil, err
	}
	return data, nil
}
func Set(path string, data []byte) error {
	if path == "" {
		if !gjson.GetBytes(bufferBytes, "@this").Exists() {
			err := errors.New("not exist")
			return err
		}
		bufferBytes = data
	}
	if path != "" {
		if !gjson.GetBytes(bufferBytes, path).Exists() {
			err := errors.New("not exist")
			return err
		}
		var err error
		bufferBytes, err = sjson.SetRawBytes(bufferBytes, path, data)
		if err != nil {
			return err
		}
	}
	return nil
}
func Del(path string) error {
	if path == "" {
		if !gjson.GetBytes(bufferBytes, "@this").Exists() {
			err := errors.New("not exist")
			return err
		}
		bufferBytes = nil
	}
	if path != "" {
		if !gjson.GetBytes(bufferBytes, path).Exists() {
			err := errors.New("not exist")
			return err
		}
		var err error
		bufferBytes, err = sjson.DeleteBytes(bufferBytes, path)
		if err != nil {
			return err
		}
	}
	return nil
}
func Pick(path string) gjson.Result {
	data := gjson.GetBytes(bufferBytes, path)
	return data
}
