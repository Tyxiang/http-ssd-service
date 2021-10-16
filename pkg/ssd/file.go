package ssd

import (
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func Load(name string) (err error) {
	path := Dir + name + ".json"
	bufferBytes, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	// if !gjson.Valid(string(bufferBytes)) {
	// 	err = errors.New("failed to load " + name)
	// 	return err
	// }
	return nil
}
func Save() (err error) {
	name := time.Now().Format("2006-01-02T15-04-05")
	path := Dir + name + ".json"
	err = ioutil.WriteFile(path, bufferBytes, 0666)
	if err != nil {
		return err
	}
	return nil
}
func Read(name string) (data []byte, err error) {
	path := Dir + name + ".json"
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func List() (names []string, err error) {
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		name := strings.TrimRight(file.Name(), ".json")
		names = append(names, name)
	}
	return names, nil
}
func Remove(name string) (err error) {
	path := Dir + name + ".json"
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
