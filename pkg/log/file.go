package log

import (
	"io/ioutil"
	"os"
	"strings"
)

func Read(name string) ([]byte, error) {
	path := Dir + name + ".log"
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
		name = strings.TrimRight(file.Name(), ".log")
		names = append(names, name)
	}
	return names, nil
}
func Remove(name string) error {
	path := Dir + name + ".log"
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
