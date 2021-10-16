package config

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

func Load(path string) (err error) {
	bufferBytes, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	// config file must be json
	if !gjson.Valid(string(bufferBytes)) {
		err = errors.New("config file format error")
		return err
	}
	return nil
}
func Save(path string) (err error) {
	err = ioutil.WriteFile(path, bufferBytes, 0666)
	if err != nil {
		return err
	}
	return nil
}
func List(dir string, ext string) (names []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		name := strings.TrimRight(file.Name(), ext) // .json
		names = append(names, name)
	}
	return names, nil
}
func New(path string, data []byte) (err error) {
	err = ioutil.WriteFile(path, data, 0644) //存在就覆盖；不存在创建。
	if err != nil {
		return err
	}
	return nil
}
func Remove(path string) (err error) {
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
