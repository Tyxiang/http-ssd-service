package log

import (
	"io/ioutil"
	"os"
	"strings"
)

var dir = "logs/"

func Read(name string) (data []byte, err error) {
	path := dir + name + ".log"
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func List() (names []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		name := strings.TrimRight(file.Name(), ".log")
		names = append(names, name)
	}
	return names, nil
}
func Remove(name string) (err error) {
	path := dir + name + ".log"
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
