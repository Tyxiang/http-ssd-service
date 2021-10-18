package ssd

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func Load(name string) error {
	path := Dir + name + ".json"
	var err error
	bufferBytes, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	// if !gjson.Valid(string(bufferBytes)) {
	// 	err := errors.New("failed to load " + name)
	// 	return err
	// }
	return nil
}
func Save() error {
	name := time.Now().Format("2006-01-02T15-04-05")
	path := Dir + name + ".json"
	err := ioutil.WriteFile(path, bufferBytes, 0666)
	if err != nil {
		return err
	}
	return nil
}
func Read(name string) ([]byte, error) {
	path := Dir + name + ".json"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func List() ([]string, error) {
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		return nil, err
	}
	name := ""
	names := []string{}
	for _, file := range files {
		name = strings.TrimRight(file.Name(), ".json")
		names = append(names, name)
	}
	return names, err
}
func Remove(name string) error {
	path := Dir + name + ".json"
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
