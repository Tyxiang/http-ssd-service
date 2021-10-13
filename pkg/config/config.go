package config

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var bufferBytes []byte
var dir = "configs/"

//// file
func Load(name string) (err error) {
	path := dir + name + ".json"
	bufferBytes, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if !gjson.Valid(string(bufferBytes)) {
		err = errors.New("failed to load " + name)
		return err
	}
	return nil
}
func Save() (err error) {
	//name := time.Now().Format("2006-01-02T15-04-05")
	name := "last"
	path := dir + name + ".json"
	err = ioutil.WriteFile(path, bufferBytes, 0666)
	if err != nil {
		return err
	}
	return nil
}
func List() (names []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		names = append(names, strings.TrimRight(file.Name(), ".json"))
	}
	return names, nil
}
func New(name string, data []byte) (err error) {
	path := dir + name + ".json"
	err = ioutil.WriteFile(path, data, 0644) //存在就覆盖；不存在创建。
	if err != nil {
		return err
	}
	return nil
}
func Remove(name string) (err error) {
	path := dir + name + ".json"
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

//// data
func Add(path string, data []byte) (err error) {
	if path == "" {
		if gjson.GetBytes(bufferBytes, "@this").Exists() {
			err = errors.New("already exist")
			return err
		}
		bufferBytes = data
	}
	if path != "" {
		if gjson.GetBytes(bufferBytes, path).Exists() {
			err = errors.New("already exist")
			return err
		}
		bufferBytes, err = sjson.SetRawBytes(bufferBytes, path, data)
		if err != nil {
			return err
		}
	}
	return nil
}
func Get(path string) (data interface{}, err error) {
	if path == "" {
		path = "@this"
	}
	data = gjson.GetBytes(bufferBytes, path).Value()
	if data == nil {
		err = errors.New("not exist")
		return nil, err
	}
	return data, nil
}
func Set(path string, data []byte) (err error) {
	if path == "" {
		if !gjson.GetBytes(bufferBytes, "@this").Exists() {
			err = errors.New("not exist")
			return err
		}
		bufferBytes = data
	}
	if path != "" {
		if !gjson.GetBytes(bufferBytes, path).Exists() {
			err = errors.New("not exist")
			return err
		}
		bufferBytes, err = sjson.SetRawBytes(bufferBytes, path, data)
		if err != nil {
			return err
		}
	}
	return nil
}
func Del(path string) (err error) {
	if path == "" {
		if !gjson.GetBytes(bufferBytes, "@this").Exists() {
			err = errors.New("not exist")
			return err
		}
		bufferBytes = nil
	}
	if path != "" {
		if !gjson.GetBytes(bufferBytes, path).Exists() {
			err = errors.New("not exist")
			return err
		}
		bufferBytes, err = sjson.DeleteBytes(bufferBytes, path)
		if err != nil {
			return err
		}
	}
	return nil
}
