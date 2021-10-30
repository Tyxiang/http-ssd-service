package config

import (
	"errors"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func Add(path string, data interface{}) error {
	if path == "" {
		if gjson.Get(buffer, "@this").Exists() {
			err := errors.New("already exist")
			return err
		}
		buffer = data.(string)
	}
	if path != "" {
		if gjson.Get(buffer, path).Exists() {
			err := errors.New("already exist")
			return err
		}
		var err error
		buffer, err = sjson.Set(buffer, path, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func Get(path string) interface{} {
	if path == "" {
		path = "@this"
	}
	data := gjson.Get(buffer, path).Value()
	return data
}

func Set(path string, data interface{}) error {
	if path == "" {
		if !gjson.Get(buffer, "@this").Exists() {
			err := errors.New("not exist")
			return err
		}
		buffer = data.(string)
	}
	if path != "" {
		if !gjson.Get(buffer, path).Exists() {
			err := errors.New("not exist")
			return err
		}
		var err error
		buffer, err = sjson.Set(buffer, path, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func Del(path string) error {
	if path == "" {
		if !gjson.Get(buffer, "@this").Exists() {
			err := errors.New("not exist")
			return err
		}
		buffer = ""
	}
	if path != "" {
		if !gjson.Get(buffer, path).Exists() {
			err := errors.New("not exist")
			return err
		}
		var err error
		buffer, err = sjson.Delete(buffer, path)
		if err != nil {
			return err
		}
	}
	return nil
}
